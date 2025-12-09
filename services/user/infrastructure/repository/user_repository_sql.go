package repository

import (
	"context"
	"database/sql"
	"errors"
	"services/user/domain"
	"services/user/repository"
)

type UserRepositorySQL struct {
	db *sql.DB
}

func NewUserRepositorySQL(db *sql.DB) repository.UserRepository {
	return &UserRepositorySQL{db: db}
}

func (r *UserRepositorySQL) Create(ctx context.Context, u *domain.User) error {
	query := `INSERT INTO users.user (id, created_by, created_at, updated_by, updated_at, sub, email, full_name, username)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)`
	_, err := r.db.ExecContext(ctx, query, u.ID, u.CreatedBy, u.CreatedAt, u.UpdatedBy, u.UpdatedAt, u.Sub, u.Email, u.Name, u.Username)
	return err
}

func (r *UserRepositorySQL) GetBySub(ctx context.Context, sub string) (*domain.User, error) {
	var u domain.User
	query := `SELECT id, created_by, created_at, updated_by, updated_at, sub, email, full_name, username
			  FROM users.user WHERE sub = $1`
	err := r.db.QueryRowContext(ctx, query, sub).Scan(&u.ID, &u.CreatedBy, &u.CreatedAt, &u.UpdatedBy, &u.UpdatedAt, &u.Sub, &u.Email, &u.Name, &u.Username)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	return &u, nil
}
