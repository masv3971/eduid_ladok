package httpserver

import (
	"eduid_ladok/internal/apiv1"
)

// Apiv1 interface
type Apiv1 interface {
	LadokInfo(indata *apiv1.RequestLadokInfo) (*apiv1.ReplyLadokInfo, error)
	SchoolInfo(indata *apiv1.RequestSchoolInfo) (*apiv1.ReplySchoolInfo, error)
}
