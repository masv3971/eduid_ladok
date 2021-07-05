package eduidiam

import (
	"context"
	"eduid_ladok/pkg/logger"
	"net/http"
	"sync"
	"time"
)

// Config holds the configuration for eduidiam
type Config struct {

	// EduIDIAMAPIURL points to eduidiam api endpoint
	EduIDIAMAPIURL string `envconfig:"EDUID_IAM_API_URL" required:"true" split_words:"true"`
}

// Service holds object for eduidiam
type Service struct {
	config     Config
	logger     *logger.Logger
	wg         *sync.WaitGroup
	httpClient *http.Client
}

// New creates a new Service object for eduidiam
func New(ctx context.Context, config Config, wg *sync.WaitGroup, logger *logger.Logger) (*Service, error) {
	s := &Service{
		config: config,
		logger: logger,
		wg:     wg,
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	return s, nil
}

// Close closes eduIDIAM
func (s *Service) Close(ctx context.Context) error {
	s.logger.Warn("Quit")

	ctx.Done()

	return nil
}
