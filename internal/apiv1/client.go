package apiv1

import (
	"context"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/trace"
)

// Client holds the publicapi object
type Client struct {
	config      *model.Cfg
	logger      *logger.Logger
	ladoks      map[string]*ladok.Service
	schoolNames []string
	tp          trace.Tracer
}

// New creates a new instance of publicapi
func New(ctx context.Context, config *model.Cfg, ladoks map[string]*ladok.Service, logger *logger.Logger) (*Client, error) {
	c := &Client{
		config:      config,
		logger:      logger,
		ladoks:      ladoks,
		schoolNames: []string{"kf", "lnu"}, // TODO(masv): "fix this!"
		tp:          otel.Tracer("Apiv1"),
	}
	_, span := c.tp.Start(ctx, "apiv1.New")
	defer span.End()

	c.logger.Info("Started")

	return c, nil
}
