package app

import (
	"github.com/caarlos0/env/v6"
	"github.com/fairytale5571/privat_test/pkg/database"
	"github.com/fairytale5571/privat_test/pkg/logger"
	"github.com/fairytale5571/privat_test/pkg/models"
	"github.com/fairytale5571/privat_test/pkg/router"
	"github.com/fairytale5571/privat_test/pkg/storage/redis"
)

type App struct {
	Logger *logger.Wrapper
	DB     *database.DB
}

func New() (*App, error) {
	log := logger.New("app")

	cfg := models.Config{}
	if err := env.Parse(&cfg); err != nil {
		log.Errorf("error parse config: %v", err)
		return nil, err
	}

	db, err := database.New(cfg.PostgresURI())
	if err != nil {
		log.Errorf("error open database: %v", err)
		return nil, err
	}

	rdb, err := redis.New(cfg.Redis)
	if err != nil {
		log.Errorf("error open redis: %v", err)
		return nil, err
	}

	rout := router.New(cfg, db, rdb)
	go rout.Start()

	return &App{
		Logger: log,
		DB:     db,
	}, nil
}
