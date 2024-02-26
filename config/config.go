package config

type Config struct {
	YandexAPIKey  string `env:"YANDEX_APIKEY,default=''"`
	YandexEnabled bool   `env:"YANDEX_ENABLED,default=false"`
	DadataEnabled bool   `env:"DADATA_ENABLED,default=false"`
	RedisSentinel bool   `env:"REDIS_SENTINEL,default=false"`

	// Deprecated: Use RedisSentinelList instead
	RedisSentinelHost string `env:"REDIS_SENTINEL_HOST,default='127.0.0.1'"`
	// Deprecated: Use RedisSentinelList instead
	RedisSentinelPort string `env:"REDIS_SENTINEL_PORT,default=26379"`

	RedisSentinelName string `env:"REDIS_SENTINEL_NAME,default='mymaster'"`
	RedisSentinelList string `env:"REDIS_SENTINEL_LIST"`

	RedisHost string `env:"REDIS_HOST,default=127.0.0.1"`
	RedisPort string `env:"REDIS_PORT,default=6379"`
	RedisAuth string `env:"REDIS_AUTH,default=''"`
	RedisDB   int    `env:"REDIS_DB,default=0"`
	AppPort   string `env:"APP_PORT,default=8080"`
	AppMode   string `env:"APP_MODE,default=prod"`
}
