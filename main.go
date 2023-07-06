package main

import (
	"address-suggesstion-proxy/config"
	"address-suggesstion-proxy/internal/controllers"
	"address-suggesstion-proxy/internal/providers"
	"address-suggesstion-proxy/internal/reposirories"
	"address-suggesstion-proxy/internal/services"
	"fmt"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/redis/go-redis/v9"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	})
	log.SetLevel(log.DebugLevel)

	cfg, err := config.LoadConfig()

	if err != nil {
		log.WithError(err).Fatal("failed to load configuration file")
	}

	db := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%s", cfg.RedisHost, cfg.RedisPort),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	repo, _ := reposirories.NewRedisRepository(db)
	cache := services.NewCacheService(repo)
	prov := providers.NewProvidersContainer(cfg)

	app := iris.Default()

	app.Party("/api").ConfigureContainer(func(r *router.APIContainer) {
		r.RegisterDependency(cache)
		r.RegisterDependency(prov)

		r.Get("/search", controllers.SearchByText)
	})

	if cfg.AppMode == "prod" {
		app.Configure(iris.WithOptimizations)
	}

	app.Listen(fmt.Sprintf(":%s", cfg.AppPort))
}
