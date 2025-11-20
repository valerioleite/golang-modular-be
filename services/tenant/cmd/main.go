package main

import (
	"context"
	dbLib "libraries/db"
	httpLib "libraries/http"
	"log/slog"
	"os"
	tenantHttp "services/tenant/internal/tenant/delivery/http"
	"services/tenant/internal/tenant/infrastructure/client"
	"services/tenant/internal/tenant/infrastructure/repository"
	"services/tenant/internal/tenant/service"

	"github.com/joho/godotenv"
)

func main() {
	setupLogger()
	setupEnvFile()

	database := setupDatabase()
	defer closeDatabaseConnection(database)

	runMigrations(database)
	injectDependencies(database)
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

func injectDependencies(database *dbLib.DB) {
	storageClient := client.NewStorageClient()
	storageRepo := repository.NewStorageRepositoryHttp(storageClient)
	tenantRepo := repository.NewTenantRepositorySQL(database.DB)
	tenantService := service.NewTenantService(tenantRepo, storageRepo)
	router := tenantHttp.NewRouter(tenantService)
	httpServer := httpLib.NewServer(router)

	startHttpServer(httpServer)
}

func startHttpServer(httpServer *httpLib.Server) {
	if err := httpServer.Start(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
