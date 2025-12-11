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
	query := `INSERT INTO users.user (id, created_by, created_at, updated_by, updated_at, sub, email, username, first_name, last_name)
			  VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)`
	_, err := r.db.ExecContext(ctx, query, u.ID, u.CreatedBy, u.CreatedAt, u.UpdatedBy, u.UpdatedAt, u.Sub, u.Email, u.Username, u.FirstName, u.LastName)
	return err
}

func (r *UserRepositorySQL) GetBySub(ctx context.Context, sub string) (*domain.User, error) {
	var u domain.User
	var username, firstName, lastName sql.NullString
	query := `SELECT id, created_by, created_at, updated_by, updated_at, sub, email, username, first_name, last_name
			  FROM users.user WHERE sub = $1`
	err := r.db.QueryRowContext(ctx, query, sub).Scan(&u.ID, &u.CreatedBy, &u.CreatedAt, &u.UpdatedBy, &u.UpdatedAt, &u.Sub, &u.Email, &username, &firstName, &lastName)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}

	if username.Valid {
		u.Username = &username.String
	}
	if firstName.Valid {
		u.FirstName = &firstName.String
	}
	if lastName.Valid {
		u.LastName = &lastName.String
	}

	return &u, nil
}
