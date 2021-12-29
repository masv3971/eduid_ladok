package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"

	"sync"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/trace"
)

// Service holds service object for ladok
type Service struct {
	config     *model.Cfg
	wg         *sync.WaitGroup
	logger     *logger.Logger
	schoolName string
	SchoolID   int
	tp         trace.Tracer

	Certificate *CertificateService
	Atom        *AtomService
	Rest        *RestService
}

// New creates a new instance of ladok Service object
func New(ctx context.Context, config *model.Cfg, wg *sync.WaitGroup, schoolName string, channel chan *model.LadokToAggregateMSG, logger *logger.Logger) (*Service, error) {
	s := &Service{
		config:     config,
		logger:     logger,
		schoolName: schoolName,
		wg:         wg,
		tp:         otel.Tracer("Ladok"),
	}

	ctx, span := s.tp.Start(ctx, "ladok.New")
	span.SetAttributes(attribute.String("SchoolName", schoolName))
	defer span.End()

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

	if err := s.getSchoolID(ctx); err != nil {
		s.logger.Warn("getSchool", err.Error())
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
