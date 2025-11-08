package service

import (
	"context"
	"services/tenant/internal/tenant/domain"
)

type mockTenantRepository struct {
	createFunc  func(ctx context.Context, tenant *domain.Tenant) error
	getAll      func(ctx context.Context) ([]*domain.Tenant, error)
	getByIdFunc func(ctx context.Context, id string) (*domain.Tenant, error)
	updateFunc  func(ctx context.Context, tenant *domain.Tenant) error
	deleteFunc  func(ctx context.Context, id string) error
}

func (m *mockTenantRepository) Create(ctx context.Context, tenant *domain.Tenant) error {
	if m.createFunc != nil {
		return m.createFunc(ctx, tenant)
	}

	return nil
}

func (m *mockTenantRepository) GetAll(ctx context.Context) ([]*domain.Tenant, error) {
	if m.getAll != nil {
		return m.getAll(ctx)
	}

	return nil, nil
}

func (m *mockTenantRepository) GetByID(ctx context.Context, id string) (*domain.Tenant, error) {
	if m.getByIdFunc != nil {
		return m.getByIdFunc(ctx, id)
	}

	return nil, nil
}

func (m *mockTenantRepository) Update(ctx context.Context, tenant *domain.Tenant) error {
	if m.updateFunc != nil {
		return m.updateFunc(ctx, tenant)
	}

	return nil
}

func (m *mockTenantRepository) Delete(ctx context.Context, id string) error {
	if m.deleteFunc != nil {
		return m.deleteFunc(ctx, id)
	}

	return nil
}
