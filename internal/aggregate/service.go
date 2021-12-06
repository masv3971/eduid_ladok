package aggregate

import (
	"context"
	"eduid_ladok/internal/ladok"
	"eduid_ladok/pkg/logger"
	"sync"

	"github.com/masv3971/goeduidiam"
)

// Config holds the configuration for aggregate
type Config struct {
	EduIDIAMURL string `required:"true" envconfig:"EDUID_IAM_URL"`
	JWTURL      string `required:"true" envconfig:"JWT_URL"`
	SamlName    string `required:"true" envconfig:"SAML_NAME"`
}

// Service holds the service object for aggregate
type Service struct {
	config      Config
	logger      *logger.Logger
	wg          *sync.WaitGroup
	ladok       *ladok.Service
	eduidiam    *goeduidiam.Client
	feedName    string
	quitChannel chan bool
}

// New creates a new instance of aggregate
func New(ctx context.Context, config Config, wg *sync.WaitGroup, feedName string, ladok *ladok.Service, logger *logger.Logger) (*Service, error) {
	s := &Service{
		logger:      logger,
		config:      config,
		ladok:       ladok,
		wg:          wg,
		feedName:    feedName,
		quitChannel: make(chan bool),
	}
	s.eduidiam = goeduidiam.New(goeduidiam.Config{
		URL: s.config.EduIDIAMURL,
		Token: goeduidiam.TokenConfig{
			Certificate: []byte{},
			PrivateKey:  []byte{},
			Password:    "",
			Scope:       "",
			Type:        "",
			URL:         s.config.JWTURL,
			Key:         "",
			Client:      "",
		},
	})

	s.wg.Add(1)
	go s.run(ctx)

	s.logger.Info("Started")
	return s, nil
}

// Close closes aggregate service
func (s *Service) Close(ctx context.Context) error {
	s.quitChannel <- true
	ctx.Done()
	s.wg.Done()

	s.logger.Warn("Quit")
	return nil
}
