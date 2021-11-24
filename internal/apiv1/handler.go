package apiv1

import (
	"context"
	"eduid_ladok/pkg/model"
	"errors"
	"time"

	"github.com/masv3971/goladok3"
)

// RequestLadokInfo request
type RequestLadokInfo struct {
	SchoolName string         `uri:"schoolName" validate:"required"`
	Data       model.UserData `json:"data" validate:"required"`
}

// ReplyLadokInfo reply
type ReplyLadokInfo struct {
	ESI           string    `json:"esi"`
	IsStudent     bool      `json:"is_student"`
	ExpireStudent time.Time `json:"expire_student"`
}

// LadokInfo handler
func (c *Client) LadokInfo(indata *RequestLadokInfo) (*ReplyLadokInfo, error) {
	ladok, ok := c.ladoks[indata.SchoolName]
	if !ok {
		return nil, errors.New("Error, can't find any matching ladok instance")
	}

	reply, _, err := ladok.Rest.Ladok.Studentinformation.GetStudent(context.TODO(), &goladok3.GetStudentReq{
		Personnummer: indata.Data.NIN,
	})
	if err != nil {
		return nil, err
	}

	replyLadokInfo := &ReplyLadokInfo{
		ESI:           ESI(reply.ExterntUID),
		IsStudent:     false,
		ExpireStudent: time.Time{},
	}
	return replyLadokInfo, nil
}

// RequestSchoolNames request
type RequestSchoolNames struct{}

// ReplySchoolNames reply
type ReplySchoolNames struct {
	SchoolNames []string `json:"school_names"`
}

// SchoolNames return a list of schoolNames
func (c *Client) SchoolNames(indata *RequestSchoolNames) (*ReplySchoolNames, error) {
	replySchoolNames := &ReplySchoolNames{
		SchoolNames: c.schoolNames,
	}
	return replySchoolNames, nil
}
