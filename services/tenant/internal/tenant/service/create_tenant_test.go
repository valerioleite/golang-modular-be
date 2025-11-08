package service

import (
	"errors"
	"services/tenant/internal/tenant/domain"
	"testing"
)

func TestTenantService_Create(t *testing.T) {
	t.Run("should return error when name is empty", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Create(nil, "", nil, nil)
		if !errors.Is(err, domain.ErrNameRequired) {
			t.Errorf("expected error %v, got %v", domain.ErrNameRequired, err)
		}
	})

	t.Run("should create tenant successfully", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)
		name := "Tenant X"
		tenant, err := service.Create(nil, name, nil, nil)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
		if tenant == nil {
			t.Fatalf("expected tenant to be created, got nil")
		}

		if tenant.Name != name {
			t.Errorf("expected name %v, got %v", name, tenant.Name)
		}
	})
}
