package ladok

import (
	"eduid_ladok/pkg/model"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
)

// GetStudent return student information from ladok
func (s *RestService) GetStudent(studentUID model.StudentUID) (*model.SIStudentRest, error) {
	url := fmt.Sprintf("%s/%s", s.Service.config.LadokRestURL, studentUID)
	resp, err := s.httpClient.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		s.logger.Warn("Status error", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Warn("Read body error", err)
	}

	reply := &model.SIStudentRest{}

	xml.Unmarshal(data, reply)

	return reply, nil
}
