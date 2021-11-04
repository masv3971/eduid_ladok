package ladok

import (
	"context"
	"eduid_ladok/pkg/model"
	"time"
)

func (s *AtomService) run(ctx context.Context) {
	superFeed, _, err := s.ladok.Feed.Recent(ctx)
	if err != nil {
		s.logger.Warn("Error", err)
	}
	if superFeed.ID == "" {
		s.logger.Info("Nothing to process")
		return
	}

	if s.isCached(superFeed.ID) {
		return
	}

	for _, superEvent := range superFeed.SuperEvents {
		if s.isCached(superEvent.ID) {
			continue
		}
		s.logger.Info("Adding message to queue", superEvent.EntryID)

		channelEvent := model.LadokToAggregateMSG{
			Event: superEvent,
			TS:    time.Now(),
		}
		s.Channel <- &channelEvent
		s.addToCache(superEvent.EntryID)
	}

	if err := s.addToCache(superFeed.ID); err != nil {
		s.logger.Warn("addToCache", err)
	}
}

func (s *AtomService) isCached(id string) bool {
	found, err := s.db.HGet(context.TODO(), s.Service.schoolName, id).Bool()
	if err != nil {
		s.logger.Warn("isCached", err)
		return false
	}
	if found {
		s.logger.Info("current id already processed, resting a bit", id)
	}

	return found
}

func (s *AtomService) addToCache(id string) error {
	if err := s.db.HSet(context.TODO(), s.Service.schoolName, "latest", id).Err(); err != nil {
		return err
	}
	s.logger.Info("Adding id to cache", id)

	return nil
}
