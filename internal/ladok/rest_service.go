package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"time"

	"github.com/masv3971/goladok3"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// RestService holds the restservice
type RestService struct {
	Service *Service
	logger  *logger.Logger
	Ladok   *goladok3.Client
	tp      trace.Tracer
}

// NewRestService creates a new instance of rest
func NewRestService(ctx context.Context, service *Service, logger *logger.Logger) (*RestService, error) {
	s := &RestService{
		logger:  logger,
		Service: service,
		tp:      otel.Tracer("Rest"),
	}

	ctx, span := s.tp.Start(ctx, "NewRestService")
	defer span.End()

	var err error
	s.Ladok, err = goladok3.New(goladok3.Config{
		URL:            s.Service.config.Ladok.URL,
		ProxyURL:       s.Service.config.HTTPProxy,
		Certificate:    s.Service.Certificate.Cert,
		CertificatePEM: s.Service.Certificate.CertPEM,
		PrivateKey:     s.Service.Certificate.PrivateKey,
		PrivateKeyPEM:  s.Service.Certificate.PrivateKeyPEM,
	})
	if err != nil {
		return nil, err
	}

	if err := s.Ladok.CheckPermission(ctx, s.Service.config.Ladok.Permissions); err != nil {
		s.logger.Warn("Ladok permission", "error", err.Error())
	}

	s.logger.Info("Started")
	return s, nil
}

// StatusLadok return the status of ladok
func (s *RestService) StatusLadok(ctx context.Context) *model.Status {
	status := &model.Status{
		Name:       "Ladok rest",
		SchoolName: s.Service.schoolName,
		Healthy:    false,
		Status:     model.StatusFail,
		Timestamp:  time.Now(),
	}

	ctx, span := s.tp.Start(ctx, "rest.StatusLadok")
	span.SetAttributes(attribute.String("SchoolName", s.Service.schoolName))
	defer span.End()

	data, _, err := s.Ladok.Kataloginformation.GetGrunddataLarosatesinformation(ctx)
	if err != nil {
		status.Message = err.Error()
		return status
	}

	if data == nil {
		status.Message = "Empty return, no data"
		return status
	}

	status.Healthy = true
	status.Status = model.StatusOK

	return status
}

// Close closes serice ladok rest
func (s *RestService) Close(ctx context.Context) error {
	_, span := s.tp.Start(ctx, "rest.quit")
	span.End()

	s.logger.Info("Quit")
	ctx.Done()
	return nil
}
