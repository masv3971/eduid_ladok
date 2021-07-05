package ladok

import (
	"eduid_ladok/pkg/logger"
	"net/http"

	"github.com/gorilla/mux"
)

// RestService holds the object for rest service
type RestService struct {
	logger     *logger.Logger
	HTTPServer *http.Server
}

// NewRestService creates a new instance of ladok rest
func NewRestService(logger *logger.Logger) *RestService {
	s := &RestService{
		logger: logger,
		HTTPServer: &http.Server{
			Addr: ":8081",
		},
	}
	router := mux.NewRouter()
	s.HTTPServer.Handler = router

	router.Use()
	//http.HandleFunc("/ladok/rest/studentinformation/{studentUID}", s.handlerSIStudent)

	router.HandleFunc("/ladok/rest/studentinformation/{studentUID}/", s.handlerSIStudent).Methods("GET")

	go func() {
		err := s.HTTPServer.ListenAndServe()
		if err != nil {
			s.logger.New("http").Fatal("listen_error", err)
		}
	}()

	s.logger.Info("Start")

	return s
}
