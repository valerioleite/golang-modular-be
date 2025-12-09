package repository

import (
	"context"
	"services/authentication/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) (*domain.User, error)
	GetBySub(ctx context.Context, sub string) (*domain.User, error)
}

