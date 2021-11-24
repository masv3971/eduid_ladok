package httpserver

import (
	"eduid_ladok/internal/apiv1"

	"github.com/gin-gonic/gin"
)

func (s *Service) endpointLadokInfo(c *gin.Context) (interface{}, error) {
	request := &apiv1.RequestLadokInfo{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.LadokInfo(request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointSchoolNames(c *gin.Context) (interface{}, error) {
	request := &apiv1.RequestSchoolNames{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.SchoolNames(request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
