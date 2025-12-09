package repository

import (
	"context"
	"services/user/domain"
)

type UserRepository interface {
	Create(ctx context.Context, user *domain.User) error
	GetBySub(ctx context.Context, sub string) (*domain.User, error)
}
