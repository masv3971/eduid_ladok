package aggregate

import (
	"context"
	"eduid_ladok/pkg/model"

	"github.com/masv3971/goladok3"
	"github.com/masv3971/goladok3/ladoktypes"
	"go.opentelemetry.io/otel/attribute"
)

func (s *Service) run(ctx context.Context) {
	ctx, span := s.tp.Start(ctx, "aggregate.run")
	span.SetAttributes(attribute.String("SchoolName", s.feedName))
	defer span.End()

	s.logger.Info("start run")
	for {
		select {
		case msg := <-s.ladok.Atom.Channel:
			s.logger.Info("received message", msg.Event.EntryID, msg.Event.EventTypeName)

			reply, _, err := s.ladok.Rest.Ladok.Studentinformation.GetAktivPaLarosate(ctx, &goladok3.GetAktivPaLarosateReq{
				UID: msg.Event.StudentUID,
			})
			if err != nil {
				s.logger.Warn(err.Error())
			}
			for _, r := range reply.Studentkopplingar {
				if r.LarosateID == s.ladok.SchoolID {
					s.logger.Info("Student", r.StudentUID, "is active!")
				}
			}
			//userReply, _, err := s.eduidiam.Users.Search(ctx, &goeduidiam.SearchUsersRequest{
			//	Data: goeduidiam.SearchRequest{
			//		Schemas:    []string{"urn:ietf:params:scim:api:messages:2.0:SearchRequest"},
			//		Filter:     "",
			//		StartIndex: 0,
			//		Count:      0,
			//		Attributes: []string{},
			//	},
			//})
			if err != nil {
				s.logger.Warn(err.Error())
			}
			//s.logger.Info("Aktiv pa larosate", reply)

			//s.

			//return
		//entry.AddTimestamp("aggregate arrived")
		//s.whatToDo(entry)
		//s.logger.Info("Process event", entry.Payload.EntryID, entry.Payload.EventType)
		case <-s.quitChannel:
			return
		}
		//	default:
		//	}
	}
}

func (s *Service) whatToDo(entry *model.LadokToAggregateMSG) {
	switch entry.Event.EventTypeName {
	case ladoktypes.LokalStudentEventName:
	default:
	}
}
