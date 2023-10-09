package main

import (
	"context"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"eduid_ladok/internal/aggregate"
	"eduid_ladok/internal/apiv1"
	"eduid_ladok/internal/httpserver"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/configuration"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
)

type service interface {
	Close(ctx context.Context) error
}

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()

	var (
		log            *logger.Logger
		mainLog        *logger.Logger
		services       = make(map[string]map[string]service)
		ladokInstances = make(map[string]*ladok.Service)
	)

	cfg, err := configuration.Parse(logger.NewSimple("Configuration"))
	if err != nil {
		panic(err)
	}

	mainLog = logger.New("main", cfg.Production)
	log = logger.New("eduid_ladok", cfg.Production)

	for schoolName := range cfg.Schools {
		ladokToAggregateChan := make(chan *model.LadokToAggregateMSG, 1000)

		s := make(map[string]service)

		ladokService, err := ladok.New(ctx, cfg, wg, schoolName, ladokToAggregateChan, log.New(schoolName).New("ladok"))
		if err != nil {
			mainLog.Error("invalid cert for", schoolName)
			continue
		}
		ladokInstances[schoolName] = ladokService
		s["ladok"] = ladokService

		aggregateService, err := aggregate.New(ctx, cfg, wg, schoolName, ladokService, log.New(schoolName).New("aggregate"))
		s["aggregate"] = aggregateService
		if err != nil {
			panic(err)
		}

		services[schoolName] = s
	}

	s := make(map[string]service)

	apiv1, err := apiv1.New(ctx, cfg, ladokInstances, log.New("apiv1"))
	if err != nil {
		panic(err)
	}
	httpserver, err := httpserver.New(ctx, cfg, apiv1, log.New("httpserver"))
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
				mainLog.Warn("feedName")
			}
		}
	}

	wg.Wait() // Block here until are workers are done

	mainLog.Info("Stopped")
}
