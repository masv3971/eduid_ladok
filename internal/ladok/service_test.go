package ladok

import (
	"bytes"
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strconv"
	"sync"
	"testing"

	"github.com/go-redis/redismock/v8"
	"github.com/masv3971/goladok3"
	"github.com/masv3971/goladok3/ladokmocks"
	"github.com/masv3971/goladok3/ladoktypes"
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

	server := httptest.NewServer(mux)

	endpoints := []struct {
		url                  string
		param                string
		paramType            string
		numericFeedEndpoint  bool
		personnummerEndpoint bool
		statusCode           int
		contentType          string
		method               string
		serverReturnPayload  []byte
	}{
		{
			url:                 "/uppfoljning/feed/recent",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(100),
		},
		{
			url:                 "/handelser/feed/recent",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(100),
		},
		{
			url:                 "/uppfoljning/feed",
			paramType:           "historicalAtom",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.XMLFeedRecent,
		},
		{
			url:                 "/handelser/feed",
			paramType:           "historicalAtom",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.XMLFeedRecent,
		},
		{
			url:                 "/uppfoljning/feed/first",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(1),
		},
		{
			url:                 "/handelser/feed/first",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeAtomXML,
			method:              "GET",
			serverReturnPayload: ladokmocks.FeedXML(1),
		},
		{
			url:                 "/kataloginformation/anvandare/autentiserad",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONKataloginformationAutentiserad,
		},
		{
			url:                 "/kataloginformation/anvandarbehorighet/egna",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONKataloginformationEgna,
		},
		{
			url:                 "/kataloginformation/behorighetsprofil",
			param:               ladokmocks.BehorighetsprofilUID,
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONKataloginformationBehorighetsprofil,
		},
		{
			url:                 "/studentinformation/student",
			param:               ladokmocks.Students[0].StudentUID,
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONStudentinformationStudent,
		},
		{
			url:                 "/studentinformation/student/personnummer",
			paramType:           "personnummer",
			statusCode:          statusCode,
			contentType:         goladok3.ContentTypeKataloginformationJSON,
			method:              "GET",
			serverReturnPayload: ladokmocks.JSONStudentinformationStudent,
		},
	}

	for _, endpoint := range endpoints {
		switch endpoint.paramType {
		case "personnummer":
			for _, student := range ladokmocks.Students {
				mockGenericEndpointServer(t, mux, endpoint.contentType, endpoint.method, endpoint.url, student.Personnummer, ladokmocks.StudentJSON(student), endpoint.statusCode)
			}
		case "historicalAtom":
			for id := 1; id <= 100; id++ {
				mockGenericEndpointServer(t, mux, endpoint.contentType, endpoint.method, endpoint.url, strconv.Itoa(id), ladokmocks.FeedXML(id), endpoint.statusCode)
			}
		default:
			mockGenericEndpointServer(t, mux, endpoint.contentType, endpoint.method, endpoint.url, endpoint.param, endpoint.serverReturnPayload, endpoint.statusCode)
		}
	}
	return server
}

func TestMockEndpoints(t *testing.T) {
	service, server, _ := mockService(t, 200, t.TempDir())
	defer server.Close()

	tts := []struct {
		name    string
		want    interface{}
		fn      interface{}
		eventID int
	}{
		{
			name:    "/handelser/feed/recent",
			want:    ladokmocks.MockSuperFeed(100),
			eventID: 100,
			fn:      service.Atom.ladok.Feed.Recent,
		},
		{
			name:    "/handelser/feed/50",
			want:    ladokmocks.MockSuperFeed(50),
			eventID: 50,
			fn:      service.Atom.ladok.Feed.Historical,
		},
		{
			name:    "/handelser/feed/first",
			want:    ladokmocks.MockSuperFeed(1),
			eventID: 1,
			fn:      service.Atom.ladok.Feed.First,
		},
		{
			name: "/kataloginformation/anvandare/autentiserad",
			want: ladokmocks.MockKataloginformationAutentiserad(),
			fn:   service.Rest.Ladok.Kataloginformation.GetAnvandareAutentiserad,
		},
		{
			name: "/kataloginformation/anvandare/egna",
			want: ladokmocks.MockKataloginformationEgna(),
			fn:   service.Rest.Ladok.Kataloginformation.GetAnvandarbehorighetEgna,
		},
		{
			name: "/kataloginformation/behorighetsprofil",
			want: ladokmocks.MockKataloginformationBehorighetsprofil(),
			fn:   service.Rest.Ladok.Kataloginformation.GetBehorighetsprofil,
		},
		{
			name:    "/studentinformation/student",
			want:    ladokmocks.MockStudentinformationStudent(),
			fn:      nil,
			eventID: 0,
		},
	}

	for _, tt := range tts {
		t.Run(tt.name, func(t *testing.T) {
			switch tt.fn.(type) {
			case func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error):
				f := tt.fn.(func(context.Context) (*ladoktypes.SuperFeed, *http.Response, error))
				reply, _, err := f(context.TODO())
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context, int) (*ladoktypes.SuperFeed, *http.Response, error):
				f := tt.fn.(func(context.Context, int) (*ladoktypes.SuperFeed, *http.Response, error))
				reply, _, err := f(context.TODO(), tt.eventID)
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error):
				f := tt.fn.(func(ctx context.Context) (*ladoktypes.KataloginformationAnvandareAutentiserad, *http.Response, error))
				reply, _, err := f(context.TODO())
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			case func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error):
				f := tt.fn.(func(context.Context) (*ladoktypes.KataloginformationAnvandarbehorighetEgna, *http.Response, error))
				reply, _, err := f(context.TODO())
				if !assert.NoError(t, err) {
					t.FailNow()
				}
				assert.Equal(t, tt.want, reply)
			default:
				t.Error("ERROR: can't find a matching reflecting type")
			}
		})
	}
}

func mockService(t *testing.T, statusCode int, tempDir string) (*Service, *httptest.Server, redismock.ClientMock) {
	server := mockLadokHTTPServer(t, statusCode)

	ladokToAggregateChan := make(chan *model.LadokToAggregateMSG, 200)

	cfg := &model.Cfg{}
	cfg.Ladok.Certificate.Folder = tempDir
	cfg.Ladok.URL = server.URL

	mockCertificate(t, 0, 1000, tempDir)

	service, err := New(context.TODO(), cfg, &sync.WaitGroup{}, "testSchoolName", ladokToAggregateChan, logger.New("test", false))
	if !assert.NoError(t, err) {
		t.FailNow()
	}

	r, redisMock := redismock.NewClientMock()
	service.Atom.db = r

	return service, server, redisMock
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
