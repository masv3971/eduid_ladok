package apiv1

import (
	"context"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
)

// Client holds the public api object
type Client struct {
	config         *model.Cfg
	logger         *logger.Logger
	ladokInstances map[string]*ladok.Service
}

// New creates a new instance of the public api
func New(ctx context.Context, config *model.Cfg, ladokInstances map[string]*ladok.Service, logger *logger.Logger) (*Client, error) {
	c := &Client{
		config:         config,
		logger:         logger,
		ladokInstances: ladokInstances,
	}

	c.logger.Info("Started")

	return c, nil
}
