package repository

import (
	"context"
	"fmt"
	"services/authentication/domain"
	"services/authentication/infrastructure/client"
	"services/authentication/infrastructure/client/dto"
	"services/authentication/repository"
	"time"

	"github.com/google/uuid"
)

type UserRepositoryHttp struct {
	client *client.UserClient
}

func NewUserRepositoryHttp(userClient *client.UserClient) repository.UserRepository {
	return &UserRepositoryHttp{client: userClient}
}

func (r *UserRepositoryHttp) Create(ctx context.Context, user *domain.User) (*domain.User, error) {
	req := &dto.CreateUserRequest{
		Sub:      user.Sub,
		Email:    user.Email,
		Name:     user.Name,
		Username: user.Username,
	}

	resp, err := r.client.CreateUser(ctx, req)
	if err != nil {
		return nil, err
	}

	return r.mapResponseToDomain(resp)
}

func (r *UserRepositoryHttp) GetBySub(ctx context.Context, sub string) (*domain.User, error) {
	resp, err := r.client.GetUserBySub(ctx, sub)
	if err != nil {
		return nil, err
	}

	if resp == nil {
		return nil, nil
	}

	return r.mapResponseToDomain(resp)
}

func (r *UserRepositoryHttp) mapResponseToDomain(resp *dto.UserResponse) (*domain.User, error) {
	id, err := uuid.Parse(resp.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to parse user ID: %w", err)
	}

	createdAt, err := time.Parse("2006-01-02T15:04:05Z07:00", resp.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse created at: %w", err)
	}

	updatedAt, err := time.Parse("2006-01-02T15:04:05Z07:00", resp.UpdatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to parse updated at: %w", err)
	}

	return &domain.User{
		ID:        id,
		CreatedBy: resp.CreatedBy,
		CreatedAt: createdAt,
		UpdatedBy: resp.UpdatedBy,
		UpdatedAt: updatedAt,
		Sub:       resp.Sub,
		Email:     resp.Email,
		Name:      resp.Name,
		Username:  resp.Username,
	}, nil
}
