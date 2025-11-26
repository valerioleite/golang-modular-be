package main

import (
	"context"
	httpLib "libraries/http"
	"log/slog"
	"os"
	authHttp "services/authentication/internal/authentication/delivery/http"
	"services/authentication/internal/authentication/infrastructure/oidc"
	"services/authentication/internal/authentication/service"

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
	oidcRepo := oidc.NewOIDCRepository()
	authService := service.NewAuthenticationService(oidcRepo)
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}
	router := authHttp.NewRouter(authService, frontendURL)
	httpServer := httpLib.NewServer(router)

	ctx := context.Background()
	err := authService.Init(ctx)
	if err != nil {
		slog.Error("Failed to initialize OIDC provider", "error", err)
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
