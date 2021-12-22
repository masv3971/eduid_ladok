package httpserver

import (
	"eduid_ladok/pkg/helpers"
	"time"

	"github.com/gin-gonic/gin"
)

func (s *Service) middlewareDuration() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		c.Next()
		duration := time.Since(t)
		c.Set("duration", duration)
	}
}

func (s *Service) middlewareLogger() gin.HandlerFunc {
	log := s.logger.New("http")
	return func(c *gin.Context) {
		c.Next()
		log.Info("request", "status", c.Writer.Status(), "url", c.Request.URL.String(), "method", c.Request.Method)
	}
}

func (s *Service) middlewareCrash() gin.HandlerFunc {
	log := s.logger.New("http")
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				status := c.Writer.Status()
				log.Error("crash", "error", r, "status", status, "url", c.Request.URL.Path, "method", c.Request.Method)
				renderContent(c, 500, gin.H{"data": nil, "error": helpers.NewError("internal_server_error")})
			}
		}()
		c.Next()
	}
}
