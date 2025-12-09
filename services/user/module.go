package user

import (
	"context"
	"database/sql"
	"log/slog"
	"services/user/delivery/http"
	"services/user/infrastructure/repository"
	"services/user/service"
)

// @title User API
// @version 1.0
// @BasePath /api/user
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.

type Module struct {
	router *http.Router
}

func NewModule(ctx context.Context, db *sql.DB) (*Module, error) {
	userRepo := repository.NewUserRepositorySQL(db)
	userSvc := service.NewUserService(userRepo)

	router := http.NewRouter(userSvc)

	slog.Info("User module initialized successfully", "module", "user")

	return &Module{
		router: router,
	}, nil
}

func (m *Module) Router() *http.Router {
	return m.router
}
