package goladok3

import (
	"context"
	"fmt"
	"net/http"
)

// studentinformationService handles studentinformation
type studentinformationService struct {
	client  *Client
	service string
}

func (s *studentinformationService) acceptHeader() string {
	return ladokAcceptHeader[s.service][s.client.format]
}

// Student is ladok reply from /studentinformation/student/{studentuid}
type Student struct {
	Avliden                           bool   `json:"Avliden"`
	Efternamn                         string `json:"Efternamn"`
	ExterntUID                        string `json:"ExterntUID"`
	FelVidEtableringExternt           bool   `json:"FelVidEtableringExternt"`
	Fodelsedata                       string `json:"Fodelsedata"`
	FolkbokforingsbevakningTillOchMed string `json:"FolkbokforingsbevakningTillOchMed"`
	Fornamn                           string `json:"Fornamn"`
	KonID                             int    `json:"KonID"`
	LarosateID                        int    `json:"LarosateID"`
	Personnummer                      string `json:"Personnummer"`
	SenastAndradAv                    string `json:"SenastAndradAv"`
	SenastSparad                      string `json:"SenastSparad"`
	UID                               string `json:"Uid"`
	UnikaIdentifierare                struct {
		LarosateID        int `json:"LarosateID"`
		UnikIdentifierare []struct {
			LarosateID     int    `json:"LarosateID"`
			SenastAndradAv string `json:"SenastAndradAv"`
			SenastSparad   string `json:"SenastSparad"`
			Typ            string `json:"Typ"`
			UID            string `json:"Uid"`
			Varde          string `json:"Varde"`
			Link           []Link `json:"link"`
		} `json:"UnikIdentifierare"`
		Link []Link `json:"link"`
	} `json:"UnikaIdentifierare"`
	Link []Link `json:"link"`
}

// GenderString translate from KonID to the equal string value
func (s *Student) GenderString() string {
	switch s.KonID {
	case 1:
		return "female"
	case 2:
		return "male"
	default:
		return "n/a"
	}
}

// GetStudentReq config for GetStudent
type GetStudentReq struct {
	UID string `validate:"required,uuid"`
}

// GetStudent return student
func (s *studentinformationService) GetStudent(ctx context.Context, req *GetStudentReq) (*Student, *http.Response, error) {
	url := fmt.Sprintf("%s/%s", s.service, "student")
	reply := &Student{}
	resp, err := s.client.call(ctx, s.acceptHeader(), "GET", url, req.UID, nil, reply)
	if err != nil {
		return nil, resp, err
	}

	return reply, resp, nil
}
