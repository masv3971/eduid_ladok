package ladok

import (
	"context"
	"errors"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
)

func (s *Service) getSchoolID(ctx context.Context) error {
	ctx, span := otel.Tracer("Rest").Start(ctx, "rest.getSchoolID")
	span.SetAttributes(attribute.String("SchoolName", s.schoolName))
	defer span.End()

	r, _, err := s.Rest.Ladok.Kataloginformation.GetGrunddataLarosatesinformation(ctx)
	if err != nil {
		return err
	}

	if len(r.Larosatesinformation) == 0 {
		return errors.New("Larosatesinformation is empty")
	}
	s.SchoolID = r.Larosatesinformation[0].LarosateID
	return nil
}
