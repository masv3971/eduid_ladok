package main

import (
	"context"
	"eduid_ladok/internal/aggregate"
	"eduid_ladok/internal/apiv1"
	"eduid_ladok/internal/httpserver"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"github.com/kelseyhightower/envconfig"
)

type config struct {
	SchoolNames []string `required:"true" split_words:"true"`
}

type service interface {
	Close(ctx context.Context) error
}

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()

	var (
		log      *logger.Logger
		mainLog  *logger.Logger
		services = make(map[string]map[string]service)
		ladoks   = make(map[string]*ladok.Service)
	)

	log = logger.New("eduid_ladok")
	mainLog = logger.New("main")

	var mainConfig config
	if err := envconfig.Process("", &mainConfig); err != nil {
		panic(err)
	}

	for _, schoolName := range mainConfig.SchoolNames {
		ladokToAggregateChan := make(chan *model.LadokToAggregateMSG, 1000)

		s := make(map[string]service)

		var ladokCfg ladok.Config
		if err := envconfig.Process(schoolName, &ladokCfg); err != nil {
			panic(err)
		}
		var aggregateCfg aggregate.Config
		if err := envconfig.Process(schoolName, &aggregateCfg); err != nil {
			panic(err)
		}

		ladok, err := ladok.New(ctx, ladokCfg, wg, schoolName, ladokToAggregateChan, log.New(schoolName).New("ladok"))
		ladoks[schoolName] = ladok
		s["ladok"] = ladok
		if err != nil {
			panic(err)
		}
		aggregate, err := aggregate.New(ctx, aggregateCfg, wg, schoolName, ladok, log.New(schoolName).New("aggregate"))
		s["aggregate"] = aggregate
		if err != nil {
			panic(err)
		}

		services[schoolName] = s
	}

	s := make(map[string]service)

	var httpserverCfg httpserver.Config
	if err := envconfig.Process("", &httpserverCfg); err != nil {
		panic(err)
	}
	var apiv1Cfg apiv1.Config
	if err := envconfig.Process("", &apiv1Cfg); err != nil {
		panic(err)
	}

	apiv1, err := apiv1.New(apiv1Cfg, ladoks, mainConfig.SchoolNames, log.New("apiv1"))
	if err != nil {
		panic(err)
	}
	httpserver, err := httpserver.New(httpserverCfg, apiv1, log.New("httpserver"))
	s["httpserver"] = httpserver
	if err != nil {
		panic(err)
	}
	services["core"] = s

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	mainLog.Info("HALTING SIGNAL!")

	for feedName, feeds := range services {
		for schoolName := range feeds {
			err := services[feedName][schoolName].Close(ctx)
			if err != nil {
				mainLog.Warn("feedName", feedName, "schoolName", schoolName, "error", err)
			}
		}
	}

	wg.Wait() // Block here until are workers are done

	mainLog.Info("Stopped")
}
