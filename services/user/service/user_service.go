package service

import (
	"context"
	"services/user/domain"
	"services/user/repository"
)

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) Create(ctx context.Context, createdBy, sub, email, name string, username *string) (*domain.User, error) {
	existingUser, err := s.userRepo.GetBySub(ctx, sub)
	if err != nil {
		return nil, err
	}

	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	user, err := domain.NewUser(createdBy, sub, email, name, username)
	if err != nil {
		return nil, err
	}

	if err := s.userRepo.Create(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserService) GetBySub(ctx context.Context, sub string) (*domain.User, error) {
	user, err := s.userRepo.GetBySub(ctx, sub)
	if err != nil {
		return nil, err
	}

	if user == nil {
		return nil, domain.ErrUserNotFound
	}

	return user, nil
}
