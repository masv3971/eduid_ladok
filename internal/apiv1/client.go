package apiv1

import (
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
)

// Client holds the publicapi object
type Client struct {
	config      *model.Cfg
	logger      *logger.Logger
	ladoks      map[string]*ladok.Service
	schoolNames []string
}

// New creates a new instance of publicapi
func New(config *model.Cfg, ladoks map[string]*ladok.Service, logger *logger.Logger) (*Client, error) {
	c := &Client{
		config:      config,
		logger:      logger,
		ladoks:      ladoks,
		schoolNames: []string{"kf", "lnu"},
	}

	c.logger.Info("Started")

	return c, nil
}
