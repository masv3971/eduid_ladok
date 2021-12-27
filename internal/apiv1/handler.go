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
	ESI             string    `json:"esi"`
	LadokExterntUID string    `json:"ladok_externt_uid"`
	IsStudent       bool      `json:"is_student"`
	ExpireStudent   time.Time `json:"expire_student"`
}

// LadokInfo handler
func (c *Client) LadokInfo(ctx context.Context, indata *RequestLadokInfo) (*ReplyLadokInfo, error) {
	ladok, ok := c.ladoks[indata.SchoolName]
	if !ok {
		return nil, errors.New("Error, can't find any matching ladok instance")
	}

	reply, _, err := ladok.Rest.Ladok.Studentinformation.GetStudent(ctx, &goladok3.GetStudentReq{
		Personnummer: indata.Data.NIN,
	})
	if err != nil {
		return nil, err
	}

	replyLadokInfo := &ReplyLadokInfo{
		ESI:             ESI(reply.ExterntUID),
		LadokExterntUID: reply.ExterntUID,
		IsStudent:       false,
		ExpireStudent:   time.Time{},
	}
	return replyLadokInfo, nil
}

// RequestSchoolInfo request
type RequestSchoolInfo struct{}

// ReplySchoolInfo reply
type ReplySchoolInfo struct {
	Schools map[string]model.SchoolInfo `json:"school_names"`
}

// SchoolInfo return a list of schoolNames
func (c *Client) SchoolInfo(ctx context.Context, indata *RequestSchoolInfo) (*ReplySchoolInfo, error) {
	replySchoolNames := &ReplySchoolInfo{}
	sn := make(map[string]model.SchoolInfo)

	for schoolName := range c.config.Schools {
		schoolInfo, ok := model.Schools[schoolName]
		if ok {
			sn[schoolName] = schoolInfo
		}
	}
	replySchoolNames.Schools = sn
	return replySchoolNames, nil
}

// Status return status for each ladok instance
func (c *Client) Status(ctx context.Context) (*model.Status, error) {
	manyStatus := model.ManyStatus{}

	for _, ladok := range c.ladoks {
		redis := ladok.Atom.StatusRedis(ctx)
		ladok := ladok.Rest.StatusLadok(ctx)

		manyStatus = append(manyStatus, redis)
		manyStatus = append(manyStatus, ladok)
	}
	status := manyStatus.Check()

	return status, nil
}
