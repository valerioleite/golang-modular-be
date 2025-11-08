package repository

import (
	"context"
	"services/tenant/internal/tenant/domain"
)

type TenantRepository interface {
	Create(ctx context.Context, tenant *domain.Tenant) error
	GetAll(ctx context.Context) ([]*domain.Tenant, error)
	GetByID(ctx context.Context, id string) (*domain.Tenant, error)
	Update(ctx context.Context, tenant *domain.Tenant) error
	Delete(ctx context.Context, id string) error
}
