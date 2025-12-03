package db

import (
	"database/sql"
	"fmt"
	"log/slog"
	"os"
	"strconv"

	_ "github.com/lib/pq"
)

type Config struct {
	Host     string
	Port     int
	User     string
	Password string
	Database string
	SSLMode  string
}

type DB struct {
	*sql.DB
}

func (d *DB) Close() error {
	if d.DB == nil {
		return nil
	}
	return d.DB.Close()
}

func NewConfig(host string, port int, user, password, database, sslMode string) *Config {
	return &Config{
		Host:     host,
		Port:     port,
		User:     user,
		Password: password,
		Database: database,
		SSLMode:  sslMode,
	}
}

func NewConfigFromEnvironment() *Config {
	databaseHost := os.Getenv("DATABASE_HOST")
	if databaseHost == "" {
		databaseHost = "localhost"
	}

	databasePort, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		databasePort = 5432
	}

	databaseUser := os.Getenv("DATABASE_USER")
	if databaseUser == "" {
		databaseUser = "postgres"
	}

	databasePassword := os.Getenv("DATABASE_PASSWORD")
	if databasePassword == "" {
		databasePassword = "postgres"
	}

	databaseDb := os.Getenv("DATABASE_DB")
	if databaseDb == "" {
		databaseDb = "postgres"
	}

	databaseSSLMode, err := strconv.ParseBool(os.Getenv("DATABASE_SSL_MODE"))
	if err != nil {
		databaseSSLMode = false
	}

	var sslMode string
	if databaseSSLMode {
		sslMode = "require"
	} else {
		sslMode = "disable"
	}

	return NewConfig(databaseHost, databasePort, databaseUser, databasePassword, databaseDb, sslMode)
}

func (c *Config) dataSourceName() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=%s",
		c.User, c.Password, c.Host, c.Port, c.Database, c.SSLMode)
}

func (c *Config) Connect() (*DB, error) {
	slog.Info("Connecting to database.", "dsn", c.dataSourceName())
	db, err := sql.Open("postgres", c.dataSourceName())
	if err != nil {
		return nil, fmt.Errorf("failed to open database: %w", err)
	}

	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping database: %w", err)
	}

	return &DB{DB: db}, nil
}
