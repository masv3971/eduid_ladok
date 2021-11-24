package httpserver

import (
	"eduid_ladok/internal/apiv1"
)

// Apiv1 interface
type Apiv1 interface {
	LadokInfo(indata *apiv1.RequestLadokInfo) (*apiv1.ReplyLadokInfo, error)
	SchoolNames(indata *apiv1.RequestSchoolNames) (*apiv1.ReplySchoolNames, error)
}
