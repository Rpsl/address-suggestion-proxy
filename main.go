package main

import (
	"address-suggesstion-proxy/config"
	"address-suggesstion-proxy/internal/controllers"
	"address-suggesstion-proxy/internal/providers"
	"address-suggesstion-proxy/internal/reposirories"
	"address-suggesstion-proxy/internal/services"
	"context"
	"fmt"
	"github.com/joho/godotenv"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/core/router"
	"github.com/sethvargo/go-envconfig"
	log "github.com/sirupsen/logrus"
	"time"
)

func main() {
	log.SetFormatter(&log.TextFormatter{
		TimestampFormat: time.RFC3339,
		FullTimestamp:   true,
	})
	log.SetLevel(log.DebugLevel)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var cfg config.Config
	if err := envconfig.Process(context.Background(), &cfg); err != nil {
		log.WithError(err).Fatal("failed to load configuration")
	}

	db := services.NewRedisClient(&cfg)
	repo, _ := reposirories.NewRedisRepository(db)
	cache := services.NewCacheService(repo)
	prov := providers.NewProvidersContainer(&cfg)

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
