package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"net/http"

	retryhttp "github.com/hashicorp/go-retryablehttp"
)

// RestService holds the restservice
type RestService struct {
	Service    *Service
	logger     *logger.Logger
	httpClient *retryhttp.Client
}

// NewRestService creates a new instance of rest
func NewRestService(ctx context.Context, service *Service, logger *logger.Logger) (*RestService, error) {
	tls, err := tlsConfig(service)
	if err != nil {
		return nil, err
	}

	s := &RestService{
		logger:  logger,
		Service: service,
		httpClient: &retryhttp.Client{
			RetryWaitMin: 1,
			RetryWaitMax: 30,
			RetryMax:     5,
			HTTPClient: &http.Client{
				Transport: &http.Transport{
					TLSClientConfig: tls,
				},
			},
		},
	}

	return s, nil
}

// Close closes serice ladok rest
func (s *RestService) Close(ctx context.Context) error {
	s.logger.Warn("Quit")

	return nil
}
