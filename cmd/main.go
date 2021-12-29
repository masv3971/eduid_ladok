package main

import (
	"context"
	"eduid_ladok/internal/aggregate"
	"eduid_ladok/internal/apiv1"
	"eduid_ladok/internal/httpserver"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/configuration"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"eduid_ladok/pkg/tracer"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"go.opentelemetry.io/otel"
)

type service interface {
	Close(ctx context.Context) error
}

func main() {
	wg := &sync.WaitGroup{}

	var (
		log      *logger.Logger
		mainLog  *logger.Logger
		services = make(map[string]map[string]service)
		ladoks   = make(map[string]*ladok.Service)
	)

	cfg, err := configuration.Parse(logger.NewSimple("Configuration"))
	if err != nil {
		panic(err)
	}

	mainLog = logger.New("main", cfg.Production)
	log = logger.New("eduid_ladok", cfg.Production)

	tp, err := tracer.New(cfg, log.New("tracer"))
	if err != nil {
		panic(err)
	}

	otel.SetTracerProvider(tp)

	// no defer span.End() since we need the span to close before main is does because main will wait until os.Signal 1 is sent.
	ctx, span := otel.Tracer("main").Start(context.Background(), "main.main")
	fmt.Println("context value", ctx.Value("transactionID"))

	for schoolName := range cfg.Schools {
		ladokToAggregateChan := make(chan *model.LadokToAggregateMSG, 1000)

		s := make(map[string]service)

		ladok, err := ladok.New(ctx, cfg, wg, schoolName, ladokToAggregateChan, log.New(schoolName).New("ladok"))
		ladoks[schoolName] = ladok
		s["ladok"] = ladok
		if err != nil {
			panic(err)
		}
		aggregate, err := aggregate.New(ctx, cfg, wg, schoolName, ladok, log.New(schoolName).New("aggregate"))
		s["aggregate"] = aggregate
		if err != nil {
			panic(err)
		}

		services[schoolName] = s
	}

	s := make(map[string]service)

	apiv1, err := apiv1.New(ctx, cfg, ladoks, log.New("apiv1"))
	if err != nil {
		panic(err)
	}
	httpserver, err := httpserver.New(ctx, cfg, apiv1, log.New("httpserver"))
	s["httpserver"] = httpserver
	if err != nil {
		panic(err)
	}
	services["core"] = s

	span.End()

	// Handle sigterm and await termChan signal
	termChan := make(chan os.Signal, 1)
	signal.Notify(termChan, syscall.SIGINT, syscall.SIGTERM)

	<-termChan // Blocks here until interrupted

	mainLog.Info("HALTING SIGNAL!")

	tp.Close(ctx)

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
