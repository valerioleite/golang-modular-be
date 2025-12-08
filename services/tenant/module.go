package tenant

import (
	"context"
	"database/sql"
	"log/slog"
	"services/tenant/delivery/http"
	"services/tenant/infrastructure/client"
	"services/tenant/infrastructure/repository"
	"services/tenant/service"
)

// @title Tenant API
// @version 1.0
// @BasePath /api/tenant
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

type Module struct {
	router *http.Router
}

func NewModule(ctx context.Context, db *sql.DB) (*Module, error) {
	storageClient := client.NewStorageClient()
	storageRepo := repository.NewStorageRepositoryHttp(storageClient)

	tenantRepo := repository.NewTenantRepositorySQL(db)
	tenantSvc := service.NewTenantService(tenantRepo, storageRepo)

	router := http.NewRouter(tenantSvc)

	slog.Info("Tenant module initialized successfully", "module", "tenant")

	return &Module{
		router: router,
	}, nil
}

func (m *Module) Router() *http.Router {
	return m.router
}
