package ladok

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"eduid_ladok/pkg/logger"
	"fmt"
	"time"
)

// CertificateService holds the certificate object
type CertificateService struct {
	Service           *Service
	logger            *logger.Logger
	CRT               *x509.Certificate
	CRTPEM            []byte
	PrivateKey        *rsa.PrivateKey
	PrivateKeyPEM     []byte
	Chain             []*x509.Certificate
	SHA256Fingerprint string
}

// NewCertificateService creates a new instance of certificate
func NewCertificateService(ctx context.Context, service *Service, logger *logger.Logger) (*CertificateService, error) {
	s := &CertificateService{
		logger:  logger,
		Service: service,
	}
	if err := s.UnwrapBundle(); err != nil {
		return nil, err
	}

	s.SHA256Fingerprint = s.NewSHA256Fingerprint()

	go func() {
		for true {
			msg := fmt.Sprintf("Certificate %q is initilized", s.SHA256Fingerprint)
			s.logger.Info(msg)

			status, notAfter := s.CheckValidTime()
			if status == Cert90DaysWarning {
				msg := fmt.Sprintf("Certificate %q expiration warning %q", s.SHA256Fingerprint, notAfter)
				s.logger.Warn(msg)
			}
			time.Sleep(24 * time.Hour)
		}
	}()

	return s, nil
}
