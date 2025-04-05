package main

import (
	"log"

	"go.uber.org/zap"

	"github.com/cyansnbrst/gesture-guru/sso-service/config"
	"github.com/cyansnbrst/gesture-guru/sso-service/internal/app"
	"github.com/cyansnbrst/gesture-guru/sso-service/pkg/db/postgres"
)

func main() {
	log.Println("starting admin-service server")

	cfg, err := config.LoadConfig("config/config-local.yml")
	if err != nil {
		log.Fatalf("loadConfig: %v", err)
	}

	logger, err := zap.NewProduction()
	if err != nil {
		log.Fatalf("can't initialize zap logger: %v", err)
	}
	defer func() {
		if err := logger.Sync(); err != nil {
			log.Printf("failed to sync logger: %v", err)
		}
	}()

	psqlDB, err := postgres.OpenDB(cfg)
	if err != nil {
		logger.Fatal("failed to init storage", zap.String("error", err.Error()))
	}
	defer psqlDB.Close()
	logger.Info("database connected")

	s := app.NewServer(cfg, logger, psqlDB)
	if err = s.Run(); err != nil {
		logger.Fatal("an error occurred", zap.String("error", err.Error()))
	}
}
