package main

import (
	"context"
	dbLib "libraries/db"
	httpLib "libraries/http"
	"log/slog"
	"net/http"
	"os"
	authHttp "services/monolith/internal/authentication/delivery/http"
	"services/monolith/internal/authentication/infrastructure/oidc"
	authService "services/monolith/internal/authentication/service"
	storageHttp "services/monolith/internal/storage/delivery/http"
	"services/monolith/internal/storage/infrastructure/s3"
	storageService "services/monolith/internal/storage/service"
	tenantHttp "services/monolith/internal/tenant/delivery/http"
	"services/monolith/internal/tenant/infrastructure/client"
	"services/monolith/internal/tenant/infrastructure/repository"
	tenantService "services/monolith/internal/tenant/service"

	"github.com/joho/godotenv"
)

func main() {
	setupLogger()
	setupEnvFile()

	database := setupDatabase()
	defer closeDatabaseConnection(database)
	runMigrations(database)

	httpServer := injectDependencies(database)
	startHttpServer(httpServer)
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

func setupDatabase() *dbLib.DB {
	config := dbLib.NewConfigFromEnvironment()
	database, err := config.Connect()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	slog.Info("Database connection established successfully.")
	return database
}

func closeDatabaseConnection(database *dbLib.DB) {
	err := database.Close()
	if err == nil {
		return
	}

	slog.Error("Failed to close database connection", "error", err)
}

func runMigrations(database *dbLib.DB) {
	migrator := dbLib.NewMigrator(database.DB, migrations, "resources/migrations")
	if err := migrator.Run(context.Background()); err != nil {
		slog.Error("Failed to run migrations", "error", err)
		os.Exit(1)
	}

	slog.Info("Migrations executed successfully.")
}

func injectDependencies(database *dbLib.DB) *httpLib.Server {
	oidcRepo := oidc.NewOIDCRepository()
	authSvc := authService.NewAuthenticationService(oidcRepo)
	frontendURL := os.Getenv("FRONTEND_URL")
	if frontendURL == "" {
		frontendURL = "http://localhost:3000"
	}

	storageRepo := s3.NewStorageRepositoryS3()
	storageSvc := storageService.NewStorageService(storageRepo)

	storageClient := client.NewStorageClient()
	storageRepoHttp := repository.NewStorageRepositoryHttp(storageClient)
	tenantRepo := repository.NewTenantRepositorySQL(database.DB)
	tenantSvc := tenantService.NewTenantService(tenantRepo, storageRepoHttp)

	ctx := context.Background()
	err := authSvc.Init(ctx)
	if err != nil {
		slog.Error("Failed to initialize OIDC provider", "error", err)
		os.Exit(1)
	}

	err = storageSvc.Init()
	if err != nil {
		slog.Error("Failed to load AWS config", "error", err)
		os.Exit(1)
	}

	authRouter := authHttp.NewRouter(authSvc, frontendURL)
	tenantRouter := tenantHttp.NewRouter(tenantSvc)
	storageRouter := storageHttp.NewRouter(storageSvc)
	router := NewRouter(authRouter, tenantRouter, storageRouter)

	return httpLib.NewServer(router)
}

func handlerFound(mux *http.ServeMux, req *http.Request) bool {
	_, pattern := mux.Handler(req)
	return pattern != ""
}

func startHttpServer(httpServer *httpLib.Server) {
	if err := httpServer.Start(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
