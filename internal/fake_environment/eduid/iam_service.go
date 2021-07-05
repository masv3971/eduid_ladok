package eduid

import "net/http"

type iamService struct {
	HTTPServer *http.Server
}

func newIAMService() *iamService {
	s := &iamService{
		HTTPServer: &http.Server{
			Addr: ":8081",
		},
	}

	return s
}
