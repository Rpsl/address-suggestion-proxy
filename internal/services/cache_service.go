package services

import (
	"address-suggesstion-proxy/internal/reposirories"
	"time"
)

type CacheService interface {
	GetByKey(key string) (string, error)
	SetByKey(key string, value string) error
	DeleteByKey(key string) error
}

type cacheService struct {
	repo reposirories.CacheRepository
}

func NewCacheService(repo reposirories.CacheRepository) CacheService {
	return &cacheService{
		repo: repo,
	}
}

func (s *cacheService) GetByKey(key string) (string, error) {
	res, _ := s.repo.Get(key)

	return res, nil
}

func (s *cacheService) SetByKey(key string, value string) error {
	err := s.repo.Set(key, value, time.Hour*24*10)

	return err
}

func (s *cacheService) DeleteByKey(key string) error {
	err := s.repo.Delete(key)

	return err
}
