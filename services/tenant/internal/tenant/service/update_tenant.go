package service

import (
	"context"
	"services/tenant/internal/tenant/domain"
)

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
