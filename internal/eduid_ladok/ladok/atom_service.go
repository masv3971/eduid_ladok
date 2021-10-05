package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"net/http"
	"time"

	memCach "github.com/patrickmn/go-cache"
)

// AtomService holds the service object
type AtomService struct {
	Service    *Service
	db         *memCach.Cache
	logger     *logger.Logger
	Channel    chan *model.ChannelEvent
	httpClient *http.Client
}

// NewAtomService creats a new instance of ladok rest
func NewAtomService(ctx context.Context, service *Service, channel chan *model.ChannelEvent, logger *logger.Logger) *AtomService {
	s := &AtomService{
		Channel: channel,
		Service: service,
		logger:  logger,
		db:      memCach.New(3*time.Hour, 6*time.Hour),
		httpClient: &http.Client{
			Timeout: 10 * time.Second,
		},
	}

	for _, atomEndpoint := range s.Service.config.LadokAtomEndpoints {
		s.Service.wg.Add(1)
		go s.run(ctx, atomEndpoint, s.logger.New(atomEndpoint))
	}
	return s
}

// Close closes ladok atom service
func (s *AtomService) Close(ctx context.Context) error {
	s.logger.Warn("Quit")
	defer func() {
		for i := 1; i <= len(s.Service.config.LadokAtomEndpoints); i++ {
			s.Service.wg.Done()
		}
	}()

	ctx.Done()
	return nil
}
