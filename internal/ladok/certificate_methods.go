package ladok

import (
	"context"
	"crypto/x509"
	"eduid_ladok/pkg/model"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"

	"go.step.sm/crypto/x509util"
)

// importCertificate read crt, key and pem file and put it the certificate handler.
func (s *CertificateService) importCertificate() error {
	var (
		certPath = filepath.Join(s.Service.config.LadokCertificateFolder, fmt.Sprintf("%s.%s", s.Service.schoolName, "crt"))
		keyPath  = filepath.Join(s.Service.config.LadokCertificateFolder, fmt.Sprintf("%s.%s", s.Service.schoolName, "key"))
	)

	for _, file := range []string{certPath, keyPath} {
		if _, err := os.Stat(file); err != nil {
			return err
		}
	}

	certFile, err := ioutil.ReadFile(certPath)
	if err != nil {
		return err
	}

	s.CertPEM = certFile

	pemBlock, _ := pem.Decode(certFile)

	s.Cert, err = x509.ParseCertificate(pemBlock.Bytes)
	if err != nil {
		return err
	}

	keyFile, err := ioutil.ReadFile(keyPath)
	if err != nil {
		return err
	}
	s.PrivateKeyPEM = keyFile

	keyBlock, _ := pem.Decode(keyFile)
	s.PrivateKey, err = x509.ParsePKCS1PrivateKey(keyBlock.Bytes)
	if err != nil {
		fmt.Println("parse privatekey")
		return err
	}

	if s.isCertificateInvalid() {
		return model.ErrCertificateNotValid
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
	if s.Cert.NotAfter.Before(days90) {
		return Cert90DaysWarning, s.Cert.NotAfter
	}

	return CertOK, s.Cert.NotAfter
}

// NewSHA256Fingerprint return fingerprint from a *x509.Certificate certificate
func (s *CertificateService) NewSHA256Fingerprint() string { return x509util.Fingerprint(s.Cert) }

func (s *CertificateService) isCertificateInvalid() bool {
	NAfterA := time.Now().After(s.Cert.NotAfter)
	NBeforeB := time.Now().Before(s.Cert.NotBefore)
	BAfterA := s.Cert.NotBefore.After(s.Cert.NotAfter)
	ABeforeB := s.Cert.NotAfter.Before(s.Cert.NotBefore)

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
