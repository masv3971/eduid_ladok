package ladok

import (
	"context"
	"eduid_ladok/pkg/logger"
	"eduid_ladok/pkg/model"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/masv3971/goladok3"
)

// AtomService holds the service object
type AtomService struct {
	Service *Service
	logger  *logger.Logger
	Channel chan *model.LadokToAggregateMSG
	ladok   *goladok3.Client
	db      *redis.Client
}

// NewAtomService create a new instance of ladok rest
func NewAtomService(ctx context.Context, service *Service, channel chan *model.LadokToAggregateMSG, logger *logger.Logger) (*AtomService, error) {
	s := &AtomService{
		Channel: channel,
		Service: service,
		logger:  logger,
	}

	switch len(s.Service.config.RedisAddr) {
	case 1:
		s.db = redis.NewClient(&redis.Options{
			Addr: s.Service.config.RedisAddr[0],
			DB:   s.Service.config.RedisDB,
		})
	default:
		s.db = redis.NewFailoverClient(&redis.FailoverOptions{
			MasterName:    "",
			SentinelAddrs: s.Service.config.RedisAddr,
			DB:            service.config.RedisDB,
		})
	}

	var err error
	s.ladok, err = goladok3.New(goladok3.Config{
		URL:            s.Service.config.LadokURL,
		Certificate:    s.Service.Certificate.Cert,
		CertificatePEM: s.Service.Certificate.CertPEM,
		PrivateKey:     s.Service.Certificate.PrivateKey,
		PrivateKeyPEM:  s.Service.Certificate.PrivateKeyPEM,
	})
	if err != nil {
		return nil, err
	}

	ticker := time.NewTicker(time.Duration(s.Service.config.LadokAtomPeriodicity) * time.Second)
	go func() {
		for {
			select {
			case <-ticker.C:
				s.run(ctx)
			case <-ctx.Done():
				s.logger.Warn("run stopped")
				ticker.Stop()
				return
			}
		}
	}()

	return s, nil
}

// Close closes ladok atom service
func (s *AtomService) Close(ctx context.Context) error {
	s.logger.Warn("Quit")

	ctx.Done()
	return nil
}
