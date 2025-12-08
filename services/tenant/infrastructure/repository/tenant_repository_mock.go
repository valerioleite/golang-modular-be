package repository

import (
	"context"
	"services/tenant/domain"
)

type TenantRepository struct {
	CreateFunc  func(ctx context.Context, tenant *domain.Tenant) error
	GetAllFunc  func(ctx context.Context) ([]*domain.Tenant, error)
	GetByIdFunc func(ctx context.Context, id string) (*domain.Tenant, error)
	UpdateFunc  func(ctx context.Context, tenant *domain.Tenant) error
	DeleteFunc  func(ctx context.Context, id string) error
}

func (m *TenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	if m.CreateFunc != nil {
		return m.CreateFunc(ctx, tenant)
	}

	return nil
}

func (m *TenantRepository) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	if m.GetAllFunc != nil {
		return m.GetAllFunc(ctx)
	}

	return nil, nil
}

func (m *TenantRepository) GetByID(ctx context.Context, id string) (*domain.Tenant, error) {
	if m.GetByIdFunc != nil {
		return m.GetByIdFunc(ctx, id)
	}

	return nil, nil
}

func (m *TenantRepository) Update(ctx context.Context, tenant *domain.Tenant) error {
	if m.UpdateFunc != nil {
		return m.UpdateFunc(ctx, tenant)
	}

	return nil
}

func (m *TenantRepository) Delete(ctx context.Context, id string) error {
	if m.DeleteFunc != nil {
		return m.DeleteFunc(ctx, id)
	}

	return nil
}
