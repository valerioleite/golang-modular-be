package storage

import (
	"context"
	"log/slog"
	"services/storage/delivery/http"
	"services/storage/infrastructure/s3"
	"services/storage/service"
)

// @title Storage API
// @version 1.0
// @BasePath /api/storage
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

type Module struct {
	router *http.Router
}

func NewModule(ctx context.Context) (*Module, error) {
	storageRepo := s3.NewStorageRepositoryS3()
	storageSvc := service.NewStorageService(storageRepo)

	if err := storageSvc.Init(); err != nil {
		return nil, err
	}

	router := http.NewRouter(storageSvc)

	slog.Info("Storage module initialized successfully")

	return &Module{
		router: router,
	}, nil
}

func (m *Module) Router() *http.Router {
	return m.router
}
