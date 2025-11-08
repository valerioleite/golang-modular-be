package service

import (
	"context"
	"strings"
	"services/tenant/internal/tenant/domain"
)

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
