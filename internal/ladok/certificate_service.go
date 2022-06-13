package ladok

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"eduid_ladok/pkg/logger"
	"fmt"
	"time"

	"go.opentelemetry.io/otel/trace"
)

// CertificateService holds the certificate object
type CertificateService struct {
	Service           *Service
	logger            *logger.Logger
	Cert              *x509.Certificate
	CertPEM           []byte
	PrivateKey        *rsa.PrivateKey
	PrivateKeyPEM     []byte
	ChainPEM          []byte
	SHA256Fingerprint string
	Chain             []*x509.Certificate
	Pkcs12            []byte
	tp                trace.Tracer
}

// NewCertificateService creates a new instance of certificate
func NewCertificateService(ctx context.Context, service *Service, logger *logger.Logger) (*CertificateService, error) {
	s := &CertificateService{
		logger:  logger,
		Service: service,
	}

	if err := s.importCertificate(ctx); err != nil {
		return nil, err
	}

	s.SHA256Fingerprint = s.NewSHA256Fingerprint(ctx)

	go func() {
		for true {
			msg := fmt.Sprintf("Certificate %q is initialized", s.SHA256Fingerprint)
			s.logger.Info(msg)

			status, notAfter := s.CheckValidTime(ctx)
			if status == Cert90DaysWarning {
				msg := fmt.Sprintf("Certificate %q expiration warning %q", s.SHA256Fingerprint, notAfter)
				s.logger.Warn(msg)
			}
			time.Sleep(24 * time.Hour)
		}
	}()

	s.logger.Info("Started")

	return s, nil
}

// Close closes certificate
func (s *CertificateService) Close(ctx context.Context) error {
	s.logger.Info("Quit")
	ctx.Done()
	return nil
}
