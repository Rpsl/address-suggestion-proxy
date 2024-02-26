package services

import (
	"address-suggesstion-proxy/config"
	"fmt"
	"strings"

	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
)

func getSentinelsList(cfg *config.Config) []string {
	if cfg.RedisSentinelList != "" {
		return strings.Split(cfg.RedisSentinelList, ",")
	}

	log.Warnln("Using deprecated sentinels config, use REDIS_SENTINEL_LIST instead")
	return []string{fmt.Sprintf("%s:%s", cfg.RedisSentinelHost, cfg.RedisSentinelPort)}
}

func NewRedisClient(cfg *config.Config) *redis.Client {
	addr := fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort)

	if cfg.RedisSentinel {
		log.Debugln("Using redis-sentinel")

		return redis.NewFailoverClient(&redis.FailoverOptions{
			SentinelAddrs: getSentinelsList(cfg),
			MasterName:    cfg.RedisSentinelName,
			Password:      cfg.RedisAuth,
			DB:            cfg.RedisDB,
		})
	}

	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: cfg.RedisAuth,
		DB:       cfg.RedisDB,
	})
}
