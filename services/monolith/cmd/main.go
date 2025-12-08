package main

import (
	"context"
	dbLib "libraries/db"
	httpLib "libraries/http"
	"log/slog"
	"net/http"
	"os"
	"services/authentication"
	"services/storage"
	"services/tenant"

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
	_ = godotenv.Load()
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

	slog.Info("Database connection established successfully")
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

	slog.Info("Migrations executed successfully")
}

func injectDependencies(database *dbLib.DB) *httpLib.Server {
	ctx := context.Background()

	authModule, err := authentication.NewModule(ctx)
	if err != nil {
		slog.Error("Failed to initialize authentication module", "error", err)
		os.Exit(1)
	}

	storageModule, err := storage.NewModule(ctx)
	if err != nil {
		slog.Error("Failed to initialize storage module", "error", err)
		os.Exit(1)
	}

	tenantModule, err := tenant.NewModule(ctx, database.DB)
	if err != nil {
		slog.Error("Failed to initialize tenant module", "error", err)
		os.Exit(1)
	}

	router := NewRouter(authModule.Router(), tenantModule.Router(), storageModule.Router())

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
