package service

import (
	"context"
	"services/tenant/internal/tenant/domain"

	"github.com/google/uuid"
)

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
