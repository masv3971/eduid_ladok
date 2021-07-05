package httpserver

import (
	"eduid_ladok/pkg/logger"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Config configure httpservice
type Config struct{}

// Service is the service object for httpserver
type Service struct {
	config Config
	logger *logger.Logger
	server *http.Server
	apiv1  Apiv1
	gin    *gin.Engine
}

// New creates a new httpserver service
func New(config *Config, apiv1 *apiv1.Client, logger *logger.Logger) (*Service, error) {
	s := &Service{
		conifg: config,
		logger: logger,
		apiv1:  apiv1,
	}

	if s.config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	
	return s, nil
}
