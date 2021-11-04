package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"

	"github.com/masv3971/goladok3"
)

// RestService holds the restservice
type RestService struct {
	Service *Service
	logger  *logger.Logger
	ladok   *goladok3.Client
}

// NewRestService creates a new instance of rest
func NewRestService(ctx context.Context, service *Service, logger *logger.Logger) (*RestService, error) {
	s := &RestService{
		logger:  logger,
		Service: service,
	}

	var err error
	s.ladok, err = goladok3.New(goladok3.Config{
		Password: s.Service.config.LadokCertificatePassword,
		//Format:       "json",
		URL:    s.Service.config.LadokURL,
		Pkcs12: s.Service.Certificate.Pkcs12,
	})
	if err != nil {
		return nil, err
	}

	return s, nil
}

// Close closes serice ladok rest
func (s *RestService) Close(ctx context.Context) error {
	s.logger.Warn("Quit")
	ctx.Done()
	return nil
}
