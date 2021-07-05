package aggregate

import (
	"context"
	"eduid_ladok/pkg/model"
)

func (s *Service) run(ctx context.Context) {
	s.logger.Info("start run")
	for {
		select {
		case entry := <-s.ladok.Atom.Channel:
			entry.AddTimestamp("aggregate arrived")
			s.whatToDo(entry)
			s.logger.Info("Process event", entry.Payload.EntryUID, entry.Payload.EventType)
		default:
		}
	}
}

func (s *Service) whatToDo(entry *model.ChannelEvent) {
	switch entry.Payload.EventType {
	case "studentinformation":
		//s.ladok.Rest.GetStudent(entry.Payload.StudentUID)
	case "studiedeltagande":
		//s.ladok.Rest.GetStudentdeltagnade(entry.Payload.StudentUID)
	}
}
