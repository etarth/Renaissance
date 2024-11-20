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
	"fmt"
)

func main() {
	cfg := config.GetConfig()
	fmt.Println('1')
	db := database.NewMongoDB(10)
	fmt.Println('1')
	s3 := s3client.NewS3Client(cfg)
	logger := logger.NewLogger(cfg)
	validator, err := validator.NewDtoValidator()
	if err != nil {
		panic(fmt.Sprintf("Failed to create dto validator: %v", err))
	}

	repositories := repositories.NewRepository(cfg, db, s3)
	usecases := usecases.NewUsecase(repositories, cfg, logger)
	handlers := handlers.NewHandler(usecases, validator)

	servers := server.NewFiberHttpServer(cfg, logger, handlers)

	servers.Start()
}
