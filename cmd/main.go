package main

import (
	"Backend/cmd/server"
	"Backend/internal/domain/usecases"
	"Backend/internal/interface/handlers"
	"Backend/internal/interface/repositories"
	"Backend/pkg/config"
	"Backend/pkg/database"
	"Backend/pkg/logger"
	"Backend/pkg/s3client"
	"Backend/pkg/validator"
	"context"
	"fmt"
	"log"
)

func main() {
	cfg := config.GetConfig()
	mongoClient, mongoDB := database.ConnectToMongoDB(cfg.GetDb().URI, cfg.GetDb().DatabaseName)
	defer func() {
		if err := mongoClient.Disconnect(context.Background()); err != nil {
			log.Fatalf("Failed to disconnect MongoDB: %v", err)
		}
	}()
	s3 := s3client.NewS3Client(cfg)
	logger := logger.NewLogger(cfg)
	validator, err := validator.NewDtoValidator()
	if err != nil {
		panic(fmt.Sprintf("Failed to create dto validator: %v", err))
	}

	repositories := repositories.NewRepository(cfg, mongoDB, s3)
	usecases := usecases.NewUsecase(repositories, cfg, logger)
	handlers := handlers.NewHandler(usecases, validator)

	servers := server.NewFiberHttpServer(cfg, logger, handlers)

	servers.Start()
}
