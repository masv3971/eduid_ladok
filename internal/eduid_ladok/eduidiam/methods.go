package eduidiam

import (
	"bytes"
	"eduid_ladok/pkg/model"
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// UpdateStudent updates a student in eduID_IAM
func (s *Service) UpdateStudent(event *model.ChannelEvent) error {
	var (
		body   model.EduIDIAMBody
		value  interface{}
		buffer io.ReadWriter
	)
	buffer = new(bytes.Buffer)
	if err := json.NewEncoder(buffer).Encode(body); err != nil {
		s.logger.Warn("error", err)
	}

	req, err := http.NewRequest("POST", s.config.EduIDIAMAPIURL, buffer)
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := s.httpClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if err := s.checkResponse(resp); err != nil {
		return err
	}

	if err := json.NewDecoder(resp.Body).Decode(value); err != nil {
		return err
	}

	return nil
}

func (s *Service) checkResponse(r *http.Response) error {
	switch r.StatusCode {
	case 200, 201, 202, 204, 304:
		return nil
	case 500:
		return errors.New("500, NotAllowed")
	}
	return errors.New("Invalid Request")
}
