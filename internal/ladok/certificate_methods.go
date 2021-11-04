package ladok

import (
	"context"
	"crypto/rsa"
	"crypto/x509"
	"eduid_ladok/pkg/model"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"go.step.sm/crypto/x509util"
	pkcs12 "software.sslmate.com/src/go-pkcs12"
)

// UnwrapBundle unwraps pfx or p12 certificate bundle into client cert, cert chain and private key
func (s *CertificateService) UnwrapBundle() error {
	certPath := filepath.Join(s.Service.config.LadokCertificateFolder, fmt.Sprintf("%s.%s", s.Service.schoolName, "pfx"))

	if _, err := os.Stat(certPath); err == nil {
		certFile, err := ioutil.ReadFile(certPath)
		if err != nil {
			return err
		}
		privateKey, clientCert, chainCerts, err := pkcs12.DecodeChain(certFile, s.Service.config.LadokCertificatePassword)
		if err != nil {
			s.logger.Warn("pkcs12 decoding", err.Error())
			return err
		}
		if s.isCertificateInvalid(clientCert) {
			return model.ErrCertificateNotValid
		}

		s.Chain = x509.NewCertPool()
		for _, chainCert := range chainCerts {
			s.Chain.AddCert(chainCert)
		}

		certPem := &pem.Block{
			Type:  "CERTIFICATE",
			Bytes: clientCert.Raw,
		}
		keyPEM := &pem.Block{
			Type:  "RSA PRIVATE KEY",
			Bytes: x509.MarshalPKCS1PrivateKey(privateKey.(*rsa.PrivateKey)),
		}

		s.CRT = clientCert
		s.PrivateKey = privateKey.(*rsa.PrivateKey)
		s.CRTPEM = pem.EncodeToMemory(certPem)
		s.PrivateKeyPEM = pem.EncodeToMemory(keyPEM)
		s.Pkcs12 = certFile
	}

	return nil
}

var (
	// Cert90DaysWarning 90 days left until not_after
	Cert90DaysWarning = "CERT_90DAYS_WARNING"
	// CertOK is ok
	CertOK = "CERT_OK"
)

// CheckValidTime keeps track of time left on certificate, return status
func (s *CertificateService) CheckValidTime() (string, time.Time) {
	days90 := time.Now().AddDate(0, 0, 90)
	if s.CRT.NotAfter.Before(days90) {
		return Cert90DaysWarning, s.CRT.NotAfter
	}

	return CertOK, s.CRT.NotAfter
}

// NewSHA256Fingerprint return fingerprint from a *x509.Certificate certificate
func (s *CertificateService) NewSHA256Fingerprint() string { return x509util.Fingerprint(s.CRT) }

func (s *CertificateService) isCertificateInvalid(cert *x509.Certificate) bool {
	NAfterA := time.Now().After(cert.NotAfter)
	NBeforeB := time.Now().Before(cert.NotBefore)
	BAfterA := cert.NotBefore.After(cert.NotAfter)
	ABeforeB := cert.NotAfter.Before(cert.NotBefore)

	if NAfterA || NBeforeB || BAfterA || ABeforeB {
		return true
	}

	return false
}

// Close closes certificate
func (s *CertificateService) Close(ctx context.Context) error {
	s.logger.Warn("Quit")
	ctx.Done()
	return nil
}
