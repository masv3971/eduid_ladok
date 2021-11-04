package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"

	"sync"
)

// Config holds the configuration for ladok
type Config struct {
	// LadokURL is the url to ladok rest api
	LadokURL string `envconfig:"LADOK_URL" required:"true" split_words:"true"`
	// LadokCertificateFolder points to the certificates file on disk
	LadokCertificateFolder string `required:"true" split_words:"true" envconfig:"LADOK_CERTIFICATE_FOLDER"` // General
	// LadokCertificatePassword password for certificates
	LadokCertificatePassword string `required:"true" split_words:"true"` // Specific
}

// Service holds service object for ladok
type Service struct {
	config     Config
	wg         *sync.WaitGroup
	logger     *logger.Logger
	schoolName string

	Certificate *CertificateService
	Atom        *AtomService
	Rest        *RestService
}

// New creates a new instance of ladok Service object
func New(ctx context.Context, config Config, wg *sync.WaitGroup, schoolName string, channel chan *model.LadokToAggregateMSG, logger *logger.Logger) (*Service, error) {
	s := &Service{
		config:     config,
		logger:     logger,
		schoolName: schoolName,
		wg:         wg,
	}

	var err error

	s.Certificate, err = NewCertificateService(ctx, s, logger.New("certificate"))
	if err != nil {
		return nil, err
	}
	s.Rest, err = NewRestService(ctx, s, logger.New("rest"))
	if err != nil {
		return nil, err
	}
	s.Atom, err = NewAtomService(ctx, s, channel, logger.New("atom"))
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Close closes ladok
func (s *Service) Close(ctx context.Context) error {
	s.Atom.Close(ctx)
	s.Rest.Close(ctx)
	s.Certificate.Close(ctx)
	return nil
}
