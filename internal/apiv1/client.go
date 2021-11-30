package apiv1

import (
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
)

// Config holds the configuration for publicapi
type Config struct {
}

// Client holds the publicapi object
type Client struct {
	config      Config
	logger      *logger.Logger
	ladoks      map[string]*ladok.Service
	schoolNames []string
}

// New creates a new instance of publicapi
func New(config Config, ladoks map[string]*ladok.Service, schoolNames []string, logger *logger.Logger) (*Client, error) {
	c := &Client{
		config:      config,
		logger:      logger,
		ladoks:      ladoks,
		schoolNames: schoolNames,
	}

	c.logger.Info("Started")

	return c, nil
}
