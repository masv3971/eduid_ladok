package model

import (
	"time"

	"github.com/masv3971/goladok3"
)

// LadokToAggregateMSG is the message on the channel, ladok-atom --> aggregate
type LadokToAggregateMSG struct {
	Event *goladok3.SuperEvent
	TS    time.Time
}

// AggregateToEduIDMSG is message on the channel from Aggregate to EduID
type AggregateToEduIDMSG struct {
}
