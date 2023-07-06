package config

import (
	"github.com/BurntSushi/toml"
	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"sync"
)

type Config struct {
	YandexAPIKey  string `toml:"yandex_apikey"`
	YandexEnabled bool   `toml:"yandex_enabled"`
	DadataEnabled bool   `toml:"dadata_enabled"`
	RedisHost     string `toml:"redis_host"`
	RedisPort     string `toml:"redis_port"`
	AppPort       string `toml:"app_port"`
	AppMode       string `toml:"app_mode"`
}

const PathConfig string = "config.toml"

var once sync.Once

// LoadConfig loads TOML configuration from a file path
func LoadConfig() (*Config, error) {
	var cfg *Config

	once.Do(func() {
		config := Config{}

		_, err := toml.DecodeFile(PathConfig, &config)

		if err != nil {
			log.Fatal(errors.Wrap(err, "failed to load config file"))
		}

		cfg = &config
	})

	return cfg, nil
}
