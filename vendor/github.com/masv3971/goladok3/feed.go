package goladok3

import (
	"context"
	"fmt"
	"net/http"
	"strconv"

	"github.com/masv3971/goladok3/ladoktypes"
)

type feedService struct {
	client  *Client
	service string
}

func (s *feedService) acceptHeader() string {
	return ladokAcceptHeader[s.service]["xml"]
}

func (s *feedService) feedURL() (string, error) {
	env, err := s.client.environment()
	if err != nil {
		return "", err
	}

	switch env {
	case ladoktypes.EnvIntTestAPI:
		return "handelser/feed", nil
	default:
		return "uppfoljning/feed", nil
	}
}

func (s *feedService) atomReader(ctx context.Context, param string) (*ladoktypes.SuperFeed, *http.Response, error) {
	envURL, err := s.feedURL()
	if err != nil {
		return nil, nil, err
	}

	url := fmt.Sprintf("%s/%s", envURL, param)

	reply := &ladoktypes.Feed{}
	resp, err := s.client.call(ctx, s.acceptHeader(), http.MethodGet, url, nil, reply)
	if err != nil {
		return nil, resp, err
	}

	superFeed, err := reply.Parse()
	if err != nil {
		return nil, resp, err
	}

	return superFeed, resp, nil

}

// Recent atom feed .../feed/recent gets the most recent publiced feed
func (s *feedService) Recent(ctx context.Context) (*ladoktypes.SuperFeed, *http.Response, error) {
	superFeed, resp, err := s.atomReader(ctx, "recent")
	if err != nil {
		return nil, resp, err
	}

	return superFeed, resp, nil
}

// HistoricalReq is config for Historical endpoint
type HistoricalReq struct {
	ID int `validate:"required"`
}

// Historical atom feed .../feed/{id} gets feed of {id}
func (s *feedService) Historical(ctx context.Context, req *HistoricalReq) (*ladoktypes.SuperFeed, *http.Response, error) {
	superFeed, resp, err := s.atomReader(ctx, strconv.Itoa(req.ID))
	if err != nil {
		return nil, resp, err
	}

	return superFeed, resp, nil
}

// First atom feed .../feed/first gets the first publiced feed
func (s *feedService) First(ctx context.Context) (*ladoktypes.SuperFeed, *http.Response, error) {
	superFeed, resp, err := s.atomReader(ctx, "first")
	if err != nil {
		return nil, resp, err
	}

	return superFeed, resp, nil

}
