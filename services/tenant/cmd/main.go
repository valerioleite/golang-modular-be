package main

import (
	"context"
	"log/slog"
	"os"
	serverHttp "services/tenant/internal/server"
	tenantHttp "services/tenant/internal/tenant/delivery/http"
	"services/tenant/internal/tenant/infrastructure/db"
	"services/tenant/internal/tenant/service"

	dbLib "libraries/db"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.LevelInfo,
	}))
	slog.SetDefault(logger)

	config := dbLib.NewConfigFromEnvironment()
	database, err := config.Connect()
	if err != nil {
		slog.Error("Failed to connect to database", "error", err)
		os.Exit(1)
	}

	defer closeDatabaseConnection(database)
	slog.Info("Database connection established successfully.")

	runMigrations(database)
	injectDependencies(database)
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
	tenantRepo := db.NewTenantRepositorySQL(database.DB)
	tenantService := service.NewTenantService(tenantRepo)
	router := tenantHttp.NewRouter(tenantService)
	httpServer := serverHttp.NewServer(router)

	startHttpServer(httpServer)
}

func startHttpServer(httpServer *serverHttp.Server) {
	if err := httpServer.Start(); err != nil {
		slog.Error("Server failed to start", "error", err)
		os.Exit(1)
	}
}
