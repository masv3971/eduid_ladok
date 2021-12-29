package httpserver

import (
	"context"
	"eduid_ladok/internal/apiv1"

	"github.com/gin-gonic/gin"
)

func (s *Service) endpointLadokInfo(ctx context.Context, c *gin.Context) (interface{}, error) {
	ctx, span := s.tp.Start(ctx, "httpserver.endpointLadokInfo")
	defer span.End()

	request := &apiv1.RequestLadokInfo{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.LadokInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointSchoolInfo(ctx context.Context, c *gin.Context) (interface{}, error) {
	ctx, span := s.tp.Start(ctx, "httpserver.endpointSchoolInfo")
	defer span.End()

	request := &apiv1.RequestSchoolInfo{}
	if err := s.bindRequest(c, request); err != nil {
		return nil, err
	}
	reply, err := s.apiv1.SchoolInfo(ctx, request)
	if err != nil {
		return nil, err
	}
	return reply, nil
}

func (s *Service) endpointStatus(ctx context.Context, c *gin.Context) (interface{}, error) {
	ctx, span := s.tp.Start(ctx, "httpserver.endpointStatus")
	defer span.End()

	reply, err := s.apiv1.Status(ctx)
	if err != nil {
		return nil, err
	}
	return reply, nil
}
