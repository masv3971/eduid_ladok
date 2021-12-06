package httpserver

import (
	"context"
	"eduid_ladok/internal/apiv1"
	"eduid_ladok/pkg/helpers"
	"eduid_ladok/pkg/logger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

// Config configure httpservice
type Config struct {
	Debug bool   `envconfig:"DEBUG"`
	Host  string `envconfig:"HOST"`
}

// Service is the service object for httpserver
type Service struct {
	config Config
	logger *logger.Logger
	server *http.Server
	apiv1  Apiv1
	gin    *gin.Engine
}

// New creates a new httpserver service
func New(config Config, api *apiv1.Client, logger *logger.Logger) (*Service, error) {
	s := &Service{
		config: config,
		logger: logger,
		apiv1:  api,
		server: &http.Server{Addr: config.Host},
	}

	if s.config.Debug {
		gin.SetMode(gin.DebugMode)
	} else {
		gin.SetMode(gin.ReleaseMode)
	}

	apiValidator := validator.New()
	binding.Validator = &defaultValidator{
		Validate: apiValidator,
	}

	s.gin = gin.New()
	s.server.Handler = s.gin
	s.server.ReadTimeout = time.Second * 5
	s.server.WriteTimeout = time.Second * 30
	s.server.IdleTimeout = time.Second * 90

	// Middlewares
	s.gin.Use(s.middlewareDuration())
	s.gin.Use(s.middlewareLogger())
	s.gin.Use(s.middlewareCrash())
	s.gin.Use(s.middlewareProbes())
	s.gin.NoRoute(func(c *gin.Context) {
		c.JSON(500, gin.H{"error": "not a valid endpoint", "data": nil})
	})

	s.regEndpoint("api/v1/:schoolName/ladokinfo", "POST", s.endpointLadokInfo)
	s.regEndpoint("api/v1/schoolinfo", "GET", s.endpointSchoolInfo)

	s.regEndpoint("api/v1/status", "GET", s.endpointStatus)

	// Run http server
	go func() {
		err := s.server.ListenAndServe()
		if err != nil {
			s.logger.New("http").Fatal("listen_error", "error", err)
		}
	}()

	s.logger.Info("started")

	return s, nil
}

func (s *Service) regEndpoint(path, method string, handler func(*gin.Context) (interface{}, error)) {
	s.gin.Handle(method, path, func(c *gin.Context) {
		res, err := handler(c)
		renderContent(c, 200, gin.H{"data": res, "error": helpers.NewErrorFromError(err)})
	})
}

func renderContent(c *gin.Context, code int, data interface{}) {
	switch c.NegotiateFormat(gin.MIMEJSON, "*/*") {
	case gin.MIMEJSON:
		c.JSON(code, data)
	case "*/*": // curl
		c.JSON(code, data)
	default:
		c.JSON(406, gin.H{"data": nil, "error": helpers.NewErrorDetails("not_acceptable", "Accept header is invalid. It should be \"application/json\".")})
	}
}

// Close closing httpserver
func (s *Service) Close(ctx context.Context) error {
	s.logger.Warn("Quit")
	return nil
}
