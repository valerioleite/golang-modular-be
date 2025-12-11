package authentication

import (
	"context"
	"log/slog"
	"services/authentication/delivery/http"
	"services/authentication/infrastructure/client"
	"services/authentication/infrastructure/oidc"
	infraRepo "services/authentication/infrastructure/repository"
	"services/authentication/service"
)

// @title Authentication API
// @version 1.0
// @BasePath /api/authentication
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description JWT token (include "Bearer " prefix when sending)

type Module struct {
	router *http.Router
}

func NewModule(ctx context.Context) (*Module, error) {
	oidcRepo := oidc.NewOIDCRepository()

	userClient := client.NewUserClient()
	userRepo := infraRepo.NewUserRepositoryHttp(userClient)

	authSvc := service.NewAuthenticationService(oidcRepo, userRepo)

	if err := authSvc.Init(ctx); err != nil {
		return nil, err
	}

	router := http.NewRouter(authSvc)

	slog.Info("Authentication module initialized successfully")

	return &Module{
		router: router,
	}, nil
}

func (m *Module) Router() *http.Router {
	return m.router
}
