package main

import (
	httpLib "libraries/http"
	"log/slog"
	"os"
	storageHttp "services/storage/internal/storage/delivery/http"
	"services/storage/internal/storage/infrastructure/s3"
	"services/storage/internal/storage/service"

	"github.com/joho/godotenv"
)

func main() {
	setupLogger()
	setupEnvFile()
	injectDependencies()
}

func setupEnvFile() {
	err := godotenv.Load()
	if err != nil {
		slog.Error(".env file not found")
		os.Exit(1)
	}
}

func setupLogger() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))

	slog.SetDefault(logger)
}

func injectDependencies() {
	storageRepo := s3.NewStorageRepositoryS3()
	storageService := service.NewStorageService(storageRepo)
	router := storageHttp.NewRouter(storageService)
	httpServer := httpLib.NewServer(router)

	err := storageService.Init()
	if err != nil {
		slog.Error("Failed to load AWS config", "error", err)
		os.Exit(1)
	}

	startHttpServer(httpServer)
}

func startHttpServer(httpServer *httpLib.Server) {
	if err := httpServer.Start(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
