package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
)

// Service holds the ladok service object
type Service struct {
	logger *logger.Logger
	Atom   *AtomService
	Rest   *RestService
}

// New creates a new service for ladok
func New(ctx context.Context, logger *logger.Logger) (*Service, error) {
	s := &Service{
		logger: logger,
	}

	s.Atom = newAtomService(s.logger.New("AtomService"))
	s.Rest = NewRestService(s.logger.New("RestService"))

	return s, nil
}

// Close closes the ladok service
func (s *Service) Close(ctx context.Context) error {
	s.logger.Info("Quit")
	return nil
}
