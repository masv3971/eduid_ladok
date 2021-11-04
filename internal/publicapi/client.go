package publicapi

import "eduid_ladok/pkg/logger"

// Config holds the configuration for publicapi
type Config struct{}

// Client holds the publicapi object
type Client struct {
	config Config
	logger *logger.Logger
}

// New creates a new instance of publicapi
func New(config Config, logger *logger.Logger) (*Client, error) {
	c := &Client{
		config: config,
		logger: logger,
	}

	return c, nil
}
