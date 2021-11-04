package ladok

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"crypto/x509/pkix"
	"eduid_ladok/pkg/logger"
	"encoding/asn1"
	"errors"
	"fmt"
	"math/big"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"software.sslmate.com/src/go-pkcs12"
)

func mockNewCertificateService(t *testing.T) *CertificateService {
	service := &Service{
		config: Config{
			LadokURL:                 "",
			LadokCertificateFolder:   t.TempDir(),
			LadokCertificatePassword: "testPassword",
		},
		wg:         &sync.WaitGroup{},
		logger:     &logger.Logger{},
		schoolName: "testSchoolName",
		Atom:       &AtomService{},
		Rest:       &RestService{},
		Certificate: &CertificateService{
			Service:    &Service{},
			logger:     &logger.Logger{},
			CRT:        &x509.Certificate{},
			PrivateKey: &rsa.PrivateKey{},
		},
	}
	certService := &CertificateService{
		Service:    service,
		logger:     &logger.Logger{},
		CRT:        &x509.Certificate{},
		PrivateKey: &rsa.PrivateKey{},
	}

	return certService
}

func mockNewLadokService(t *testing.T) *Service {
	service := &Service{
		config: Config{
			LadokURL:                 "",
			LadokCertificateFolder:   t.TempDir(),
			LadokCertificatePassword: "testPassword",
		},
		wg:          &sync.WaitGroup{},
		logger:      &logger.Logger{},
		schoolName:  "testSchoolName",
		Atom:        &AtomService{},
		Rest:        &RestService{},
		Certificate: &CertificateService{},
	}

	return service
}

func mockCertificateTemplate(t *testing.T, notBefore, notAfter int) *x509.Certificate {
	certTemplate := &x509.Certificate{
		SignatureAlgorithm: x509.SHA256WithRSA,
		PublicKeyAlgorithm: x509.RSA,
		Version:            3,
		SerialNumber:       big.NewInt(2300),
		Issuer: pkix.Name{
			Country:            []string{},
			Organization:       []string{},
			OrganizationalUnit: []string{},
			Locality:           []string{},
			Province:           []string{},
			StreetAddress:      []string{},
			PostalCode:         []string{},
			SerialNumber:       "",
			CommonName:         "Ladok3 LED MIT API CA",
			Names:              []pkix.AttributeTypeAndValue{},
			ExtraNames:         []pkix.AttributeTypeAndValue{},
		},
		Subject: pkix.Name{
			Country:            []string{"SE"},
			Organization:       []string{"Ladok"},
			OrganizationalUnit: []string{"LED", "Int-test-API"},
			Locality:           []string{"Stockholm"},
			Province:           []string{},
			StreetAddress:      []string{},
			PostalCode:         []string{},
			SerialNumber:       "",
			CommonName:         "sunet@KF",
			Names:              []pkix.AttributeTypeAndValue{},
			ExtraNames:         []pkix.AttributeTypeAndValue{},
		},
		NotBefore:                   time.Now().AddDate(0, 0, notBefore),
		NotAfter:                    time.Now().AddDate(0, 0, notAfter),
		KeyUsage:                    x509.KeyUsageDataEncipherment | x509.KeyUsageKeyEncipherment | x509.KeyUsageDigitalSignature,
		Extensions:                  []pkix.Extension{},
		ExtraExtensions:             []pkix.Extension{},
		UnhandledCriticalExtensions: []asn1.ObjectIdentifier{},
		ExtKeyUsage:                 []x509.ExtKeyUsage{x509.ExtKeyUsageClientAuth, x509.ExtKeyUsageServerAuth},
		OCSPServer:                  []string{"URI:http://ca01.rome08.led.ladok.se:8080/ca/ocsp"},
	}

	return certTemplate
}

func mockNewCertificateBundle(t *testing.T, notBefore, notAfter int, folder, fileType, password string) {
	certTemplate := mockCertificateTemplate(t, notBefore, notAfter)

	privateKey, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		t.Fatal(err)
	}

	certByte, err := x509.CreateCertificate(rand.Reader, certTemplate, certTemplate, &privateKey.PublicKey, privateKey)
	if err != nil {
		t.Fatal(err)
	}

	cert, err := x509.ParseCertificate(certByte)
	if err != nil {
		t.Fatal(err)
	}

	data, err := pkcs12.Encode(rand.Reader, privateKey, cert, []*x509.Certificate{}, password)
	if err != nil {
		t.Fatal(err)
	}

	certPath := filepath.Join(folder, fmt.Sprintf("%s.%s", "testSchoolName", fileType))

	if err := os.WriteFile(certPath, data, 0777); err != nil {
		t.Fatal(err)
	}
}

func TestUnwrapBundle(t *testing.T) {
	type have struct {
		notBefore, notAfter int
		fileType            string
	}
	tts := []struct {
		name string
		have have
		want error
	}{
		{
			name: "OK p12",
			have: have{
				notBefore: 0,
				notAfter:  100,
				fileType:  "p12",
			},
			want: errors.New(""),
		},
		{
			name: "OK pfx",
			have: have{
				notBefore: 0,
				notAfter:  100,
				fileType:  "pfx",
			},
			want: errors.New(""),
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			service, _ := mockService(t, 200)

			if err := service.Certificate.UnwrapBundle(); err != nil {
				assert.EqualError(t, err, tt.want.Error())
			}

			assert.NotNil(t, service.Certificate.CRT, "client cert should not be nil")
			assert.NotNil(t, service.Certificate.PrivateKey, "privatekey should not be nil")
			assert.NotEmpty(t, service.Certificate.CRTPEM)
			assert.NotEmpty(t, service.Certificate.PrivateKeyPEM)
			assert.NotEmpty(t, service.Certificate.CRT.NotAfter, "should not be empty")
		})
	}
}

func TestCheckValidTime(t *testing.T) {
	type have struct {
		notBefore, notAfter int
	}
	tts := []struct {
		name string
		have have
		want string
	}{
		{
			name: "OK - 100 days cert",
			have: have{
				notBefore: 0,
				notAfter:  100,
			},
			want: CertOK,
		},
		{
			name: "1 day cert",
			have: have{
				notBefore: 0,
				notAfter:  1,
			},
			want: Cert90DaysWarning,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			cs := mockNewCertificateService(t)

			cs.CRT = mockCertificateTemplate(t, tt.have.notBefore, tt.have.notAfter)

			gotStatus, _ := cs.CheckValidTime()

			t.Logf("name: %q", tt.name)
			assert.Equal(t, tt.want, gotStatus, "status should be equal")
		})
	}
}

func TestIsCertificateInvalid(t *testing.T) {
	type have struct {
		notBefore, notAfter int
	}
	tts := []struct {
		name string
		have have
		want bool
	}{
		{
			name: "Cert into the future, standard case",
			have: have{
				notBefore: 0,
				notAfter:  2,
			},
			want: false,
		},
		{
			name: "Cert into the future, standard case",
			have: have{
				notBefore: 0,
				notAfter:  100,
			},
			want: false,
		},
		{
			name: "Cert in the past, old cert",
			have: have{
				notBefore: -2,
				notAfter:  -1,
			},
			want: true,
		},
		{
			name: "NotBefore is After NotAfter, wrong cert",
			have: have{
				notBefore: 2,
				notAfter:  -2,
			},
			want: true,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			cs := mockNewCertificateService(t)
			cert := mockCertificateTemplate(t, tt.have.notBefore, tt.have.notAfter)

			got := cs.isCertificateInvalid(cert)

			assert.Equal(t, tt.want, got)
		})
	}
}
