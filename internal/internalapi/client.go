package internalapi

import (
	"eduid_ladok/pkg/logger"
)

// Config holds the configuration for internalapi
type Config struct {
}

// Client holds the internalapi object
type Client struct {
	config Config
	logger *logger.Logger
}

// New creates a new instanace of internalAPI
func New(config Config, logger *logger.Logger) (*Client, error) {
	c := &Client{
		config: config,
		logger: logger,
	}

	return c, nil
}
