package main

import (
	"log"

	"go.uber.org/zap"

	"cyansnbrst/gestures-service/config"
	"cyansnbrst/gestures-service/internal/server"
	"cyansnbrst/gestures-service/pkg/db/postgres"
)

func main() {
	log.Println("starting gestures-service server")

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

	// redisClient := redis.NewRedisClient(cfg)
	// defer func() {
	// 	if err := redisClient.Close(); err != nil {
	// 		logger.Warn("failed to close redis", zap.String("error", err.Error()))
	// 	}
	// }()
	// logger.Info("redis connected")

	s := server.NewServer(cfg, logger, psqlDB)
	if err = s.Run(); err != nil {
		logger.Fatal("an error occurred", zap.String("error", err.Error()))
	}
}
