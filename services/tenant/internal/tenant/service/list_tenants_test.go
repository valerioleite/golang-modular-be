package service

import (
	"context"
	"services/tenant/internal/tenant/domain"
	"testing"
)

func TestTenantService_List(t *testing.T) {
	t.Run("should return empty list successfully", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)

		repo.getAll = func(ctx context.Context) ([]*domain.Tenant, error) {
			return []*domain.Tenant{}, nil
		}

		tenants, err := service.List(nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(tenants) != 0 {
			t.Errorf("expected empty list, got %v tenants", len(tenants))
		}
	})

	t.Run("should return list with tenants successfully", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)

		repo.getAll = func(ctx context.Context) ([]*domain.Tenant, error) {
			tenant1, _ := domain.NewTenant("Tenant 1", nil, nil)
			tenant2, _ := domain.NewTenant("Tenant 2", nil, nil)
			return []*domain.Tenant{tenant1, tenant2}, nil
		}

		tenants, err := service.List(nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if len(tenants) != 2 {
			t.Errorf("expected 2 tenants, got %v", len(tenants))
		}
	})
}
