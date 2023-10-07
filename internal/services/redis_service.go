package services

import (
	"address-suggesstion-proxy/config"
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func NewRedisClient(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	if cfg.RedisSentinel {
		log.Debugln("using redis-sentinel")

		sentinel := redis.NewSentinelClient(&redis.Options{
			Addr: fmt.Sprintf("%s:%s", cfg.RedisSentinelHost, cfg.RedisSentinelPort),
		})

		masters, err := sentinel.GetMasterAddrByName(context.Background(), cfg.RedisSentinelName).Result()

		if err != nil {
			log.WithError(err).Fatal("failed to get redis sentinel master host")
		}

		addr = fmt.Sprintf("%s:%s", masters[0], masters[1])

		log.Debugf(fmt.Sprintf("redis master from sentinel is: %s", addr))
	}

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisAuth,
		DB:       cfg.RedisDB,
	})
}
