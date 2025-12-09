package repository

import (
	"context"
	"services/authentication/domain"
	"services/authentication/repository"
	"time"

	"github.com/google/uuid"
)

type UserRepositoryMock struct {
	CreateFunc   func(ctx context.Context, user *domain.User) (*domain.User, error)
	GetBySubFunc func(ctx context.Context, sub string) (*domain.User, error)
}

func NewUserRepositoryMock() repository.UserRepository {
	return &UserRepositoryMock{}
}

func (m *UserRepositoryMock) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, user)
	}

	now := time.Now()
	return &domain.User{
		ID:        uuid.New(),
		CreatedBy: user.CreatedBy,
		CreatedAt: now,
		UpdatedBy: user.UpdatedBy,
		UpdatedAt: now,
		Sub:       user.Sub,
		Email:     user.Email,
		Name:      user.Name,
		Username:  user.Username,
	}, nil
}

func (m *UserRepositoryMock) GetBySub(ctx context.Context, sub string) (*domain.User, error) {
	if m.GetBySubFunc != nil {
		return m.GetBySubFunc(ctx, sub)
	}

	now := time.Now()
	return &domain.User{
		ID:        uuid.New(),
		CreatedBy: "system",
		CreatedAt: now,
		UpdatedBy: "system",
		UpdatedAt: now,
		Sub:       sub,
		Email:     "test@example.com",
		Name:      "Test User",
		Username:  nil,
	}, nil
}

