package ladok

import (
	"crypto/tls"
	"crypto/x509"
	"errors"
	"fmt"
)

func tlsConfig(service *Service) (*tls.Config, error) {
	keyPair, err := tls.X509KeyPair([]byte(service.Certifiate.CRTPEM), []byte(service.Certifiate.PrivateKeyPEM))
	if err != nil {
		fmt.Println("mura")
		return nil, err
	}

	if len(keyPair.Certificate) > 1 {
		return nil, errors.New("PEM file contains multiple certificates")
	}

	tlsCfg := &tls.Config{
		Certificates:                []tls.Certificate{keyPair},
		NameToCertificate:           map[string]*tls.Certificate{},
		RootCAs:                     &x509.CertPool{},
		NextProtos:                  []string{},
		ServerName:                  service.config.LadokRestURL,
		ClientAuth:                  0,
		ClientCAs:                   &x509.CertPool{},
		InsecureSkipVerify:          false,
		CipherSuites:                []uint16{},
		PreferServerCipherSuites:    false,
		SessionTicketsDisabled:      false,
		SessionTicketKey:            [32]byte{},
		ClientSessionCache:          nil,
		MinVersion:                  0,
		MaxVersion:                  0,
		CurvePreferences:            []tls.CurveID{},
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}

	return tlsCfg, nil
}
