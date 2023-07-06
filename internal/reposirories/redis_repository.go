package reposirories

import (
	"context"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

type RedisRepository struct {
	db *redis.Client
}

func NewRedisRepository(db *redis.Client) (*RedisRepository, error) {
	return &RedisRepository{db: db}, nil
}

func (r *RedisRepository) Get(key string) (string, error) {
	key = r.normalizeKey(key)

	res, err := r.db.Get(context.Background(), key).Result()

	if err == redis.Nil {
		return "", nil
	} else if err != nil {
		log.Errorln("error redis get result", err)
		return "", err
	}

	return res, nil
}

func (r *RedisRepository) Set(key string, value string, ttl time.Duration) error {
	key = r.normalizeKey(key)

	err := r.db.Set(context.Background(), key, value, ttl).Err()

	if err != nil {
		log.Errorln(errors.Wrap(err, "error while save data in redis"))
		return err
	}

	return nil
}

func (r *RedisRepository) Delete(key string) error {
	key = r.normalizeKey(key)

	return r.db.Del(context.Background(), key).Err()
}

func (r *RedisRepository) normalizeKey(query string) string {
	query = strings.ToLower(query)
	query = strings.Trim(query, " ")

	return query
}
