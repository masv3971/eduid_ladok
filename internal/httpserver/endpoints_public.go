package httpserver

import (
	"eduid_ladok/internal/publicapi"

	"github.com/gin-gonic/gin"
)

func (s *Service) endpointPublic(c *gin.Context) (interface{}, error) {
	request := &publicapi.RequestPublic{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.publicAPI.Public(request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
