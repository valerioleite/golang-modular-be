package service

import (
	"context"
	"services/tenant/internal/tenant/domain"
)

func (s *TenantService) List(ctx context.Context) ([]*domain.Tenant, error) {
	return s.repo.GetAll(ctx)
}
