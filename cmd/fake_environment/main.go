package main

import (
	"context"
	"eduid_ladok/internal/fake_environment/eduid"
	"eduid_ladok/internal/fake_environment/ladok"
	"eduid_ladok/pkg/logger"
	"os"
	"os/signal"
	"sync"
	"syscall"
)

type service interface {
	Close(ctx context.Context) error
}

func main() {
	wg := &sync.WaitGroup{}
	ctx := context.Background()

	var (
		log      *logger.Logger
		mainLog  *logger.Logger
		services = make(map[string]service)
	)

	log = logger.New("fake_enviroment")
	mainLog = logger.New("main")

	ladok, err := ladok.New(ctx, log.New("ladok"))
	services["ladok"] = ladok
	if err != nil {
		panic(err)
	}

	eduid, err := eduid.New(ctx, log.New("eduid"))
	services["eduid"] = eduid
	if err != nil {
		panic(err)
	}

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	mainLog.Info("HALTING SIGNAL!")

	for serviceName := range services {
		err := services[serviceName].Close(ctx)
		if err != nil {
			mainLog.Warn("serviceName", serviceName, "error", err)
		}
	}

	wg.Wait() // Block here until are workers are done

	mainLog.Info("Stoped")
}
