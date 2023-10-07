package config

type Config struct {
	YandexAPIKey  string `env:"YANDEX_APIKEY,default=''"`
	YandexEnabled bool   `env:"YANDEX_ENABLED,default=false"`
	DadataEnabled bool   `env:"DADATA_ENABLED,default=false"`
	RedisSentinel bool   `env:"REDIS_SENTINEL,default=false"`
	RedisHost     string `env:"REDIS_HOST,default=127.0.0.1"`
	RedisPort     string `env:"REDIS_PORT,default=6379"`
	AppPort       string `env:"APP_PORT,default=8080"`
	AppMode       string `env:"APP_MODE,default=prod"`
}
