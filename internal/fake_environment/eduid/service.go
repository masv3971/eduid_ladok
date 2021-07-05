package eduid

import (
	"context"
	"eduid_ladok/pkg/logger"
)

// Service holds the service object for eduid
type Service struct {
	logger *logger.Logger
	IAM    *iamService
}

// New creates a new instance of eduid
func New(ctx context.Context, logger *logger.Logger) (*Service, error) {
	s := &Service{
		logger: logger,
		IAM:    newIAMService(),
	}
	return s, nil
}

// Close closes service eudid
func (s *Service) Close(ctx context.Context) error {
	s.logger.Info("Quit")
	return nil
}
