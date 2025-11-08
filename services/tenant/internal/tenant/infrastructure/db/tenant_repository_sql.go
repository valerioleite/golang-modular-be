package db

import (
	"context"
	"database/sql"
	"services/tenant/internal/tenant/domain"
	"services/tenant/internal/tenant/repository"
)

type TenantRepositorySQL struct {
	db *sql.DB
}

func NewTenantRepositorySQL(db *sql.DB) repository.TenantRepository {
	return &TenantRepositorySQL{db: db}
}

func (r *TenantRepositorySQL) Create(ctx context.Context, t *domain.Tenant) error {
	query := "INSERT INTO tenant (id, name, logo, banner) VALUES ($1, $2, $3, $4)"
	_, err := r.db.ExecContext(ctx, query, t.ID, t.Name, t.Logo, t.Banner)
	return err
}

func (r *TenantRepositorySQL) GetByID(ctx context.Context, id string) (*domain.Tenant, error) {
	var t domain.Tenant
	query := "SELECT id, name, logo, banner FROM tenant WHERE id = $1"
	err := r.db.QueryRowContext(ctx, query, id).Scan(&t.ID, &t.Name, &t.Logo, &t.Banner)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &t, nil
}

func (r *TenantRepositorySQL) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	query := "SELECT id, name, logo, banner FROM tenant ORDER BY name"
	rows, err := r.db.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tenants []*domain.Tenant
	for rows.Next() {
		var t domain.Tenant
		if err := rows.Scan(&t.ID, &t.Name, &t.Logo, &t.Banner); err != nil {
			return nil, err
		}

		tenants = append(tenants, &t)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return tenants, nil
}

func (r *TenantRepositorySQL) Update(ctx context.Context, t *domain.Tenant) error {
	query := "UPDATE tenant SET name = $1, logo = $2, banner = $3 WHERE id = $4"
	_, err := r.db.ExecContext(ctx, query, t.Name, t.Logo, t.Banner, t.ID)
	return err
}

func (r *TenantRepositorySQL) Delete(ctx context.Context, id string) error {
	query := "DELETE FROM tenant WHERE id = $1"
	_, err := r.db.ExecContext(ctx, query, id)
	return err
}

func (r *TenantRepositorySQL) CreateTable(ctx context.Context) error {
	query := `CREATE TABLE IF NOT EXISTS tenant (
		id UUID PRIMARY KEY,
		name VARCHAR(255) NOT NULL,
		logo VARCHAR(255),
		banner VARCHAR(255)
	)`
	_, err := r.db.ExecContext(ctx, query)
	return err
}
