package ladok

import (
	"bytes"
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"

	"github.com/masv3971/goladok3"
	"github.com/masv3971/goladok3/testinginfra"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func mockGenericEndpointServer(t *testing.T, mux *http.ServeMux, contentType, method, url, param string, payload []byte, statusCode int) {
	mux.HandleFunc(fmt.Sprintf("%s/%s", url, param),
		func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", contentType)
			w.WriteHeader(statusCode)
			testMethod(t, r, method)
			testURL(t, r, fmt.Sprintf("%s/%s", url, param))
			w.Write(payload)
		},
	)
}
func mockLadokHTTPServer(t *testing.T, statusCode int) *httptest.Server {
	mux := http.NewServeMux()

	//	server := httptest.NewTLSServer(mux)
	server := httptest.NewServer(mux)

	endpoints := []struct {
		url         string
		param       string
		statusCode  int
		contentType string
		method      string
		payload     []byte
	}{
		{
			url:         "/uppfoljning/feed/recent",
			param:       "",
			statusCode:  statusCode,
			contentType: goladok3.ContentTypeAtomXML,
			method:      "GET",
			payload:     testinginfra.XMLFeedRecent,
		},
		{
			url:         "/handelse/feed/recent",
			param:       "",
			statusCode:  statusCode,
			contentType: goladok3.ContentTypeAtomXML,
			method:      "GET",
			payload:     testinginfra.XMLFeedRecent,
		},
		{
			url:         "/kataloginformation/anvandare/autentiserad",
			param:       "",
			statusCode:  statusCode,
			contentType: goladok3.ContentTypeKataloginformationJSON,
			method:      "GET",
			payload:     testinginfra.XMLFeedRecent,
		},
		{
			url:         "/kataloginformation/anvandare/egna",
			param:       "",
			statusCode:  statusCode,
			contentType: goladok3.ContentTypeKataloginformationJSON,
			method:      "GET",
			payload:     testinginfra.XMLFeedRecent,
		},
		{
			url:         "/kataloginformation/behorighetsprofil",
			param:       testinginfra.BehorighetsprofilUID,
			statusCode:  statusCode,
			contentType: goladok3.ContentTypeKataloginformationJSON,
			method:      "GET",
			payload:     testinginfra.XMLFeedRecent,
		},
		{
			url:         "/studentinformation/student",
			param:       testinginfra.StudentUID,
			statusCode:  statusCode,
			contentType: goladok3.ContentTypeKataloginformationJSON,
			method:      "GET",
			payload:     testinginfra.XMLFeedRecent,
		},
	}

	for _, ep := range endpoints {
		mockGenericEndpointServer(t, mux, ep.contentType, ep.method, ep.url, ep.param, ep.payload, ep.statusCode)
	}

	return server
}

func mockService(t *testing.T, statusCode int) (*Service, *httptest.Server) {
	var (
		tempDir    = t.TempDir()
		pw         = "testPassword"
		schoolName = "testSchoolName"
		fileType   = "pfx"
	)

	server := mockLadokHTTPServer(t, statusCode)

	ladokToAggregateChan := make(chan *model.LadokToAggregateMSG, 200)

	mockNewCertificateBundle(t, 0, 1000, tempDir, fileType, pw)

	cfg := Config{
		LadokURL:                 server.URL,
		LadokCertificateFolder:   tempDir,
		LadokCertificatePassword: pw,
	}
	service, err := New(context.TODO(), cfg, &sync.WaitGroup{}, schoolName, ladokToAggregateChan, logger.New("test"))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	return service, server
}

func testMethod(t *testing.T, r *http.Request, want string) {
	got := r.Method
	assert.Equal(t, want, got)
}

func testURL(t *testing.T, r *http.Request, want string) {
	got := r.RequestURI
	assert.Equal(t, want, got)
}

func testBody(t *testing.T, r *http.Request, want string) {
	buffer := new(bytes.Buffer)
	_, err := buffer.ReadFrom(r.Body)
	assert.NoError(t, err)

	got := buffer.String()
	require.JSONEq(t, want, got)
}
