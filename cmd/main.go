package main

import (
	"context"
	"eduid_ladok/internal/aggregate"
	"eduid_ladok/internal/httpserver"
	"eduid_ladok/internal/internalapi"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/internal/publicapi"
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
	)

	log = logger.New("eduid_ladok")
	mainLog = logger.New("main")

	var mainConfig config
	if err := envconfig.Process("", &mainConfig); err != nil {
		panic(err)
	}

	for _, schoolName := range mainConfig.SchoolNames {
		ladokToAggregateChan := make(chan *model.LadokToAggregateMSG, 200)

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
	var internalAPICfg internalapi.Config
	if err := envconfig.Process("", &internalAPICfg); err != nil {
		panic(err)
	}
	var publicAPICfg publicapi.Config
	if err := envconfig.Process("", &publicAPICfg); err != nil {
		panic(err)
	}

	internalAPI, err := internalapi.New(internalAPICfg, log.New("internalAPI"))
	if err != nil {
		panic(err)
	}
	publicAPI, err := publicapi.New(publicAPICfg, log.New("publicAPI"))
	if err != nil {
		panic(err)
	}
	httpserver, err := httpserver.New(httpserverCfg, internalAPI, publicAPI, log.New("httpserver"))
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

	mainLog.Info("Stoped")
}
