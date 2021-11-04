package aggregate

import (
	"context"
	"eduid_ladok/pkg/model"

	"github.com/masv3971/goladok3"
)

func (s *Service) run(ctx context.Context) {
	s.logger.Info("start run")
	for {
		select {
		//case entry := <-s.ladok.Atom.Channel:
		//entry.AddTimestamp("aggregate arrived")
		//s.whatToDo(entry)
		//s.logger.Info("Process event", entry.Payload.EntryID, entry.Payload.EventType)
		default:
		}
	}
}

func (s *Service) whatToDo(entry *model.LadokToAggregateMSG) {
	switch entry.Event.EventTypeName {
	case goladok3.LokalStudentEventName:
	case goladok3.AnvandareAndradEventName:
	}
}
