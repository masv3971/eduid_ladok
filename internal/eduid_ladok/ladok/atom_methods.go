package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"encoding/xml"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	memCach "github.com/patrickmn/go-cache"
)

func (s *AtomService) run(ctx context.Context, endpoint string, logger *logger.Logger) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			time.Sleep(1 * time.Second)
			msg, err := s.getFeed(endpoint, logger)
			if err != nil {
				time.Sleep(2 * time.Second)
			} else {
				msg.MakeChannelEvents()
				for _, event := range msg.Events {
					event.AddTimestamp("ladok_atom")
					if !s.inCach(event.Payload.EntryUID, logger.New("inCach")) {
						if err := s.addToCach(event.Payload, logger.New("addToCach")); err != nil {
							logger.Warn("error", err)
						}

						s.Channel <- event
						logger.Info("Added event to channel", "EntryUID", event.Payload.EntryUID, "studentUID", event.Payload.StudentUID)
					}
				}
			}
		}
	}
}

func (s *AtomService) getFeed(endpoint string, logger *logger.Logger) (*model.MSG, error) {
	//resp, err := s.httpClient.Get(s.Service.config.LadokAtomURL)
	url := fmt.Sprintf("%s/%s/atom/%s", s.Service.config.LadokAtomURL, s.Service.schoolName, endpoint)
	resp, err := s.httpClient.Get(url)
	if err != nil {
		logger.Warn("Get error", err)
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		logger.Warn("Status error", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Warn("Read body error", err)
	}

	feed := &model.Feed{}

	xml.Unmarshal(data, feed)

	msg := &model.MSG{
		Feed: feed,
	}
	return msg, nil
}

func (s *AtomService) inCach(id string, logger *logger.Logger) bool {
	_, found := s.db.Get(id)
	logger.Info("check if record is present in cach DB", id, found)
	return found
}

func (s *AtomService) addToCach(payload *model.EventPayload, logger *logger.Logger) error {
	logger.Info("Add new record in cach DB")
	if err := s.db.Add(payload.EntryUID, payload.StudentUID, memCach.DefaultExpiration); err != nil {
		return err
	}
	return nil
}
