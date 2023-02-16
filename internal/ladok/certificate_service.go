package ladok

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"fmt"
	"time"
)

// CertificateService holds the certificate object
type CertificateService struct {
	Service                 *Service
	logger                  *logger.Logger
	Cert                    *x509.Certificate
	CertPEM                 []byte
	PrivateKey              *rsa.PrivateKey
	PrivateKeyPEM           []byte
	ChainPEM                []byte
	SHA256Fingerprint       string
	Chain                   []*x509.Certificate
	Pkcs12                  []byte
	ClientCertificateStatus *model.MonitoringCertClient
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
			s.logger.Info(fmt.Sprintf("Certificate %q is initialized", s.SHA256Fingerprint))
			status, notAfter := s.CheckValidTime(ctx)
			clientCertificateStatus := &model.MonitoringCertClient{
				Valid:       true,
				Fingerprint: s.SHA256Fingerprint,
				NotAfter:    notAfter,
				DaysLeft:    certDaysLeft(notAfter),
				LastChecked: time.Now(),
			}

			switch status {
			case Cert90DaysWarning:
				s.logger.Warn(fmt.Sprintf("Certificate %q expiration warning %q", s.SHA256Fingerprint, notAfter))
			case CertExpired:
				s.logger.Error(fmt.Sprintf("Certificate %q is expired %q", s.SHA256Fingerprint, notAfter))
				clientCertificateStatus.Valid = false
			}

			s.ClientCertificateStatus = clientCertificateStatus
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
