package service

import (
	"context"
	"github.com/google/uuid"
	"services/tenant/internal/tenant/domain"
	"services/tenant/internal/tenant/repository"
	"strings"
)

type TenantService struct {
	repo repository.TenantRepository
}

func NewTenantService(repo repository.TenantRepository) *TenantService {
	return &TenantService{repo: repo}
}

func (s *TenantService) Create(ctx context.Context, name string, logo, banner *string) (*domain.Tenant, error) {
	if strings.TrimSpace(name) == "" {
		return nil, domain.ErrNameRequired
	}

	tenant, err := domain.NewTenant(name, logo, banner)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Create(ctx, tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}
func (s *TenantService) List(ctx context.Context) ([]*domain.Tenant, error) {
	return s.repo.GetAll(ctx)
}

func (s *TenantService) Get(ctx context.Context, id string) (*domain.Tenant, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, domain.ErrInvalidID
	}

	tenant, err := s.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tenant == nil {
		return nil, domain.ErrTenantNotFound
	}

	return tenant, nil
}

func (s *TenantService) Update(ctx context.Context, id, name string, logo, banner *string) (*domain.Tenant, error) {
	tenant, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	err = tenant.Update(name, logo, banner)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Update(ctx, tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *TenantService) Delete(ctx context.Context, id string) error {
	_, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	return s.repo.Delete(ctx, id)
}
