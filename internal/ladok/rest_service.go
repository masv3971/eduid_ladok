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
	Ladok   *goladok3.Client
}

// NewRestService creates a new instance of rest
func NewRestService(ctx context.Context, service *Service, logger *logger.Logger) (*RestService, error) {
	s := &RestService{
		logger:  logger,
		Service: service,
	}

	var err error
	s.Ladok, err = goladok3.New(goladok3.Config{
		URL:            s.Service.config.LadokURL,
		ProxyURL:       s.Service.config.HTTPProxy,
		Certificate:    s.Service.Certificate.Cert,
		CertificatePEM: s.Service.Certificate.CertPEM,
		PrivateKey:     s.Service.Certificate.PrivateKey,
		PrivateKeyPEM:  s.Service.Certificate.PrivateKeyPEM,
	})
	if err != nil {
		return nil, err
	}

	s.logger.Info("Started")
	return s, nil
}

// Close closes serice ladok rest
func (s *RestService) Close(ctx context.Context) error {
	s.logger.Warn("Quit")
	ctx.Done()
	return nil
}
