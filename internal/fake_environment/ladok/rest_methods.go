package ladok

import (
	"eduid_ladok/pkg/model"
	"encoding/xml"
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

type studentinformationRequest struct {
}
type studentinformationReply struct {
	Status bool `json:"name"`
}

func (s *RestService) handlerSIStudent(w http.ResponseWriter, r *http.Request) {
	s.logger.Info("handlerSIStudent")
	vars := mux.Vars(r)
	fmt.Println(vars["StudentUID"])

	reply := model.SIStudentRest{
		Avliden:                           false,
		Efternamn:                         "testEfternamn",
		ExterntUID:                        "testExterntUID",
		FelVidEtableringExternt:           false,
		FolkbokforingsbevakningTillOchMed: time.Time{},
		Fodelsedata:                       "",
		Fornamn:                           "",
		KonID:                             0,
		Personnummer:                      "",
	}

	w.Header().Set("Content-Type", "application/vnd.ladok+xml")
	data, err := xml.MarshalIndent(reply, "", "   ")
	if err != nil {
		s.logger.Warn("Marshal_error", err)
	}

	fmt.Fprintf(w, string(data))
}
