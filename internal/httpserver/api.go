package httpserver

import (
	"context"
	"eduid_ladok/internal/apiv1"
	"eduid_ladok/pkg/model"
)

// Apiv1 interface
type Apiv1 interface {
	LadokInfo(ctx context.Context, indata *apiv1.RequestLadokInfo) (*apiv1.ReplyLadokInfo, error)
	SchoolInfo(ctx context.Context, indata *apiv1.RequestSchoolInfo) (*apiv1.ReplySchoolInfo, error)
	Status(ctx context.Context) (*model.Status, error)
}
