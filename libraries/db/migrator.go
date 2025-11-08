package db

import (
	"context"
	"crypto/sha256"
	"database/sql"
	"embed"
	"encoding/hex"
	"fmt"
	"io/fs"
	"path/filepath"
	"sort"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

type Migrator struct {
	db         *sql.DB
	migrations embed.FS
	path       string
}

func NewMigrator(db *sql.DB, migrations embed.FS, path string) *Migrator {
	return &Migrator{
		db:         db,
		migrations: migrations,
		path:       path,
	}
}

func (m *Migrator) createMigrationsTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS schema_migrations (
		id UUID PRIMARY KEY,
		version INTEGER NOT NULL,
		filename VARCHAR(255) NOT NULL,
		checksum VARCHAR(64) NOT NULL,
		applied_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
		UNIQUE(version),
		UNIQUE(filename)
	)`
	_, err := m.db.ExecContext(ctx, query)
	return err
}

type MigrationRecord struct {
	Version  int
	Filename string
	Checksum string
}

func (m *Migrator) getAppliedMigrations(ctx context.Context) (map[int]*MigrationRecord, error) {
	if err := m.createMigrationsTable(ctx); err != nil {
		return nil, fmt.Errorf("failed to create migrations table: %w", err)
	}

	rows, err := m.db.QueryContext(ctx, "SELECT version, filename, checksum FROM schema_migrations ORDER BY version")
	if err != nil {
		return nil, fmt.Errorf("failed to query migrations: %w", err)
	}
	defer rows.Close()

	applied := make(map[int]*MigrationRecord)
	for rows.Next() {
		var record MigrationRecord
		if err := rows.Scan(&record.Version, &record.Filename, &record.Checksum); err != nil {
			return nil, fmt.Errorf("failed to scan migration record: %w", err)
		}
		applied[record.Version] = &record
	}

	return applied, rows.Err()
}

func (m *Migrator) calculateChecksum(content []byte) string {
	hash := sha256.Sum256(content)
	return hex.EncodeToString(hash[:])
}

func (m *Migrator) recordMigrationTx(ctx context.Context, tx *sql.Tx, version int, filename, checksum string) error {
	query := `INSERT INTO schema_migrations (id, version, filename, checksum)
		VALUES ($1, $2, $3, $4)
		ON CONFLICT (version) DO NOTHING`
	_, err := tx.ExecContext(ctx, query, uuid.New(), version, filename, checksum)
	return err
}

func (m *Migrator) getMigrationFiles() ([]string, error) {
	var files []string

	err := fs.WalkDir(m.migrations, m.path, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return nil
		}

		if !strings.HasSuffix(path, ".sql") {
			return nil
		}

		files = append(files, path)
		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("failed to walk migrations directory: %w", err)
	}

	sort.Strings(files)
	return files, nil
}

func (m *Migrator) extractVersion(filename string) (int, error) {
	base := filepath.Base(filename)
	parts := strings.Split(base, "_")
	if len(parts) == 0 {
		return 0, fmt.Errorf("invalid filename format: %s", filename)
	}

	version, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, fmt.Errorf("invalid version in filename %s: %w", filename, err)
	}

	return version, nil
}

func (m *Migrator) Run(ctx context.Context) error {
	applied, err := m.getAppliedMigrations(ctx)
	if err != nil {
		return fmt.Errorf("failed to get applied migrations: %w", err)
	}

	files, err := m.getMigrationFiles()
	if err != nil {
		return fmt.Errorf("failed to get migration files: %w", err)
	}

	for _, file := range files {
		version, err := m.extractVersion(file)
		if err != nil {
			return fmt.Errorf("failed to extract version from %s: %w", file, err)
		}

		filename := filepath.Base(file)

		if record, exists := applied[version]; exists {
			content, err := m.migrations.ReadFile(file)
			if err != nil {
				return fmt.Errorf("failed to read migration file %s: %w", file, err)
			}

			checksum := m.calculateChecksum(content)
			if record.Checksum != checksum {
				return fmt.Errorf("checksum mismatch for migration %s: expected %s, got %s", filename, record.Checksum, checksum)
			}
			continue
		}

		content, err := m.migrations.ReadFile(file)
		if err != nil {
			return fmt.Errorf("failed to read migration file %s: %w", file, err)
		}

		checksum := m.calculateChecksum(content)

		tx, err := m.db.BeginTx(ctx, nil)
		if err != nil {
			return fmt.Errorf("failed to begin transaction: %w", err)
		}

		if _, err := tx.ExecContext(ctx, string(content)); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("failed to execute migration %s: %w", file, err)
		}

		if err := m.recordMigrationTx(ctx, tx, version, filename, checksum); err != nil {
			_ = tx.Rollback()
			return fmt.Errorf("failed to record migration %s: %w", filename, err)
		}

		if err := tx.Commit(); err != nil {
			return fmt.Errorf("failed to commit migration %s: %w", filename, err)
		}
	}

	return nil
}
