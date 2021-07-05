package ladok

import (
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"fmt"
	"net/http"
)

// AtomService holds the service object for atom
type AtomService struct {
	logger     *logger.Logger
	HTTPServer *http.Server
}

func newAtomService(logger *logger.Logger) *AtomService {
	s := &AtomService{
		logger: logger,
		HTTPServer: &http.Server{
			Addr: ":8080",
		},
	}

	for school := range model.Schools {
		url := fmt.Sprintf("/%s/atom/studentinformation", school)
		s.logger.Info("Atom endpoint", url)
		http.HandleFunc(url, s.handlerSI)
	}

	for school := range model.Schools {
		url := fmt.Sprintf("/%s/atom/studiedeltagande", school)
		s.logger.Info("Atom endpoint", url)
		http.HandleFunc(url, s.handlerSD)
	}

	go func() {
		if err := s.HTTPServer.ListenAndServe(); err != nil {
			s.logger.New("http").Fatal("listen_error", err)
		}
	}()

	s.logger.Info("Start")

	return s
}
