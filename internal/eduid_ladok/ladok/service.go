package ladok

import (
	"context"
	"eduid_ladok/internal/eduid_ladok/eduidiam"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"

	"sync"
)

// Config holds the configuration for ladok
type Config struct {
	// LadokURL is the specefic url to ladok instance
	LadokAtomURL string `envconfig:"LADOK_ATOM_URL" required:"true" split_words:"true"`
	// LadokRestAPIURL is the url to ladok rest api
	LadokRestURL string `envconfig:"LADOK_REST_URL" required:"true" split_words:"true"`
	// LadokAtomEndpoints is a list of endpoints to fetch
	LadokAtomEndpoints []string `envconfig:"LADOK_ATOM_ENDPOINTS" required:"true"`
	// LadokCertificatePath points to the certificates file on disk
	LadokCertificatePath string `required:"true" split_words:"true" envconfig:"LADOK_CERTIFICATE_PATH"` // General
	// LadokCertificatePassword password for certificates
	LadokCertificatePassword string `required:"true" split_words:"true"` // Specefic
}

// Service holds service object for ladok
type Service struct {
	config     Config
	wg         *sync.WaitGroup
	logger     *logger.Logger
	eduid      *eduidiam.Service
	schoolName string

	Atom       *AtomService
	Rest       *RestService
	Certifiate *CertificateService
}

// New creates a new instance of ladok Service object
func New(ctx context.Context, config Config, wg *sync.WaitGroup, schoolName string, channel chan *model.ChannelEvent, logger *logger.Logger) (*Service, error) {
	s := &Service{
		config:     config,
		logger:     logger,
		schoolName: schoolName,
		wg:         wg,
	}

	var err error

	s.Certifiate, err = NewCertificateService(ctx, s, logger.New("certificate"))
	if err != nil {
		return nil, err
	}
	s.Rest, err = NewRestService(ctx, s, logger.New("rest"))
	if err != nil {
		return nil, err
	}
	s.Atom = NewAtomService(ctx, s, channel, logger.New("atom"))

	return s, nil
}

// Close closes ladok
func (s *Service) Close(ctx context.Context) error {
	s.Atom.Close(ctx)
	s.Certifiate.Close(ctx)
	return nil
}
