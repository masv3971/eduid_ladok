package goladok3

import (
	"bytes"
	"context"
	"crypto/rand"
	"crypto/rsa"
	"crypto/tls"
	"crypto/x509"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"time"

	"github.com/masv3971/goladok3/ladoktypes"
	"golang.org/x/time/rate"
)

// Config configures new function
type Config struct {
	URL            string            `validate:"required"`
	Certificate    *x509.Certificate `validate:"required"`
	CertificatePEM []byte            `validate:"required"`
	PrivateKey     *rsa.PrivateKey   `validate:"required"`
	PrivateKeyPEM  []byte            `validate:"required"`
	ProxyURL       string
	//Chain         []*x509.Certificate `validate:"required"`
}

// Client holds the ladok object
type Client struct {
	//password       string
	HTTPClient     *http.Client
	rateLimit      *rate.Limiter
	format         string
	url            string
	certificate    *x509.Certificate
	certificatePEM []byte
	chain          *x509.CertPool
	chainPEM       []byte
	privateKey     *rsa.PrivateKey
	privateKeyPEM  []byte
	proxyURL       string

	Kataloginformation *kataloginformationService
	Studentinformation *studentinformationService
	Studentdeltagande  *studentdeltagandeService
	Feed               *feedService
}

// New create a new instanace of ladok
func New(config Config) (*Client, error) {
	if err := Check(config); err != nil {
		return nil, err
	}
	c := &Client{
		format:         "json",
		url:            config.URL,
		proxyURL:       config.ProxyURL,
		privateKeyPEM:  config.PrivateKeyPEM,
		certificatePEM: config.CertificatePEM,
		certificate:    config.Certificate,
		privateKey:     config.PrivateKey,
		rateLimit:      rate.NewLimiter(rate.Every(1*time.Second), 30),
	}

	if err := c.httpConfigure(); err != nil {
		return nil, err
	}

	c.Studentinformation = &studentinformationService{client: c, service: "studentinformation"}
	c.Kataloginformation = &kataloginformationService{client: c, service: "kataloginformation"}
	c.Studentdeltagande = &studentdeltagandeService{client: c, service: "studentdeltagande"}
	c.Feed = &feedService{client: c, service: "feed"}

	return c, nil
}

func (c *Client) httpConfigure() error {
	keyPair, err := tls.X509KeyPair(c.certificatePEM, c.privateKeyPEM)
	if err != nil {
		return err
	}

	tlsCfg := &tls.Config{
		Rand:         rand.Reader,
		Certificates: []tls.Certificate{keyPair},
		NextProtos:   []string{},
		ClientAuth:   tls.RequireAndVerifyClientCert,
		//ClientCAs:          c.chainDER,
		InsecureSkipVerify: false,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
		},
		PreferServerCipherSuites:    true,
		SessionTicketsDisabled:      false,
		DynamicRecordSizingDisabled: false,
		Renegotiation:               0,
		KeyLogWriter:                nil,
	}

	tlsCfg.BuildNameToCertificate()

	//	var proxyConfig func(*http.Request) (*url.URL, error)
	//	if c.proxyURL != "" {
	//		proxyURL, err := url.Parse(c.proxyURL)
	//		if err != nil {
	//			return err
	//		}
	//
	//		proxyConfig = http.ProxyURL(proxyURL)
	//		fmt.Println("LOGGING, using proxy:", proxyURL)
	//	} else {
	//		proxyConfig = nil
	//		fmt.Println("LOGGING, no proxy is using")
	//	}

	c.HTTPClient = &http.Client{
		Transport: &http.Transport{
			TLSClientConfig:     tlsCfg,
			DialContext:         nil,
			TLSHandshakeTimeout: 30 * time.Second,
			//	Proxy:               http.ProxyFromEnvironment,
		},
	}

	if os.Getenv("DEBUG") == "true" {
		fmt.Println("DEBUG HTTPClient transport", c.HTTPClient.Transport)
	}

	return nil
}

var ladokAcceptHeader = map[string]map[string]string{
	"studentinformation": {
		"json": "application/vnd.ladok-studentinformation+json",
		"xml":  "application/vnd.ladok-studentinformation+xml",
	},
	"studentdeltagande": {
		"json": "application/vnd.ladok-studiedeltagande+json",
		"xml":  "application/vnd.ladok-studiedeltagande+xml",
	},
	"resultat": {
		"json": "application/vnd.ladok-resultat+json",
		"xml":  "application/vnd.ladok-resultat+xml",
	},
	"uppfoljning": {
		"json": "application/vnd.ladok-uppfoljning+json",
		"xml":  "application/vnd.ladok-uppfoljning+xml",
	},
	"examen": {
		"json": "application/vnd.ladok-examen+json",
		"xml":  "application/vnd.ladok-examen+xml",
	},
	"utbildningsinformation": {
		"json": "application/vnd.ladok-utbildningsinformation+json",
		"xml":  "application/vnd.ladok-utbildningsinformation+xml",
	},
	"kataloginformation": {
		"json": "application/vnd.ladok-kataloginformation+json",
		"xml":  "application/vnd.ladok-kataloginformation+xml",
	},
	"feed": {
		"xml": "application/atom+xml",
	},
}

// NewRequest make a new request
func (c *Client) newRequest(ctx context.Context, acceptHeader string, method, path string, body interface{}) (*http.Request, error) {
	rel, err := url.Parse(path)
	if err != nil {
		return nil, oneError("", "url", "newRequest", err.Error())
	}

	u, err := url.Parse(c.url)
	if err != nil {
		return nil, oneError("", "url", "newRequest", err.Error())
	}
	url := u.ResolveReference(rel)

	var buf io.ReadWriter
	if body != nil {
		payload := struct {
			Data interface{} `json:"data"`
		}{
			Data: body,
		}
		buf = new(bytes.Buffer)
		err = json.NewEncoder(buf).Encode(payload)
		if err != nil {
			return nil, oneError("", "json", "newRequest", err.Error())
		}
	}

	req, err := http.NewRequestWithContext(ctx, method, url.String(), buf)
	if err != nil {
		return nil, oneError("", "http.NewRequestWithContext", "newRequest", err.Error())
	}

	proxyURL, err := http.ProxyFromEnvironment(req)
	if err != nil {
		return nil, err
	}

	c.HTTPClient = &http.Client{
		Transport: &http.Transport{
			ProxyConnectHeader: make(http.Header),
			Proxy:              http.ProxyURL(proxyURL),
		},
	}

	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	req.Header.Set("Accept", acceptHeader)
	req.Header.Set("User-Agent", "goladok3/0.0.15")

	if os.Getenv("DEBUG") == "true" {
		out, _ := httputil.DumpRequest(req, false)
		fmt.Println("DEBUG dumprequest", string(out))
	}

	return req, nil
}

// Do does the new request
func (c *Client) do(req *http.Request, value interface{}) (*http.Response, error) {
	ctx := context.Background()
	if err := c.rateLimit.Wait(ctx); err != nil {
		return nil, oneError("", "HTTPClient.Do", "ratelimit", err.Error())
	}

	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, oneError("Can't perform http.client.do", "HTTPClient.Do", "do", err.Error())
	}
	defer resp.Body.Close()

	if os.Getenv("DEBUG") == "true" {
		out, _ := httputil.DumpResponse(resp, false)
		fmt.Println("DEBUG dumpResponse", string(out))
	}

	if err := checkResponse(resp); err != nil {
		buf := &bytes.Buffer{}
		if _, err := buf.ReadFrom(resp.Body); err != nil {
			return nil, oneError("Can't process buffer", "buf.ReadFrom", "do", err.Error())
		}
		ladokError := &ladoktypes.LadokError{}
		if err := json.Unmarshal(buf.Bytes(), ladokError); err != nil { // TODO(masv): Fix xml error parsing into Errors.
			return nil, oneError("Can't unmarshal json to Errors ", "json.Unmarshal", "do", err.Error())
		}
		return nil, &Errors{Ladok: ladokError}
	}

	switch resp.Header.Get("Content-Type") {
	case ContentTypeAtomXML:
		if err := xml.NewDecoder(resp.Body).Decode(value); err != nil {
			return nil, err
		}
	case ContentTypeKataloginformationJSON, ContentTypeStudiedeltagandeJSON, ContentTypeStudentinformationJSON:
		if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
			return nil, err
		}
	default:
		return nil, ErrNoValidContentType
	}

	return resp, nil
}

func checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return oneError("Not allowed", "statusCode", "checkResponse", "")
	}
	return oneError("Invalid request", "statusCode", "checkResponse", "")
}

func (c *Client) call(ctx context.Context, acceptHeader, method, url string, req, reply interface{}) (*http.Response, error) {
	request, err := c.newRequest(
		ctx,
		acceptHeader,
		method,
		url,
		req,
	)
	if err != nil {
		return nil, err
	}

	resp, err := c.do(request, reply)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
