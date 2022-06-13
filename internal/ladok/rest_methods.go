package ladok

import (
	"context"
	"errors"
)

func (s *Service) getSchoolID(ctx context.Context) error {
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
