package service

import (
	"context"
	"errors"
	"services/tenant/internal/tenant/domain"
	"services/tenant/internal/tenant/infrastructure/repository"
	"testing"

	"github.com/google/uuid"
)

func TestTenantService_Create(t *testing.T) {
	t.Run("should return error when name is empty", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Create(nil, "", nil, nil)
		if !errors.Is(err, domain.ErrNameRequired) {
			t.Errorf("expected error %v, got %v", domain.ErrNameRequired, err)
		}
	})

	t.Run("should create tenant successfully", func(t *testing.T) {
		repo := &repository.TenantRepository{}
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

func TestTenantService_List(t *testing.T) {
	t.Run("should return empty list successfully", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		repo.GetAllFunc = func(ctx context.Context) ([]*domain.Tenant, error) {
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
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		repo.GetAllFunc = func(ctx context.Context) ([]*domain.Tenant, error) {
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

func TestTenantService_Get(t *testing.T) {
	t.Run("should return error when id is empty", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Get(nil, "")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Get(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not registered", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		repo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		_, err := service.Get(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should get tenant successfully", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		name := "Test Tenant 123"
		repo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse("df5872a7-a907-4e5a-bc5c-23a79e595baa")
			return &domain.Tenant{ID: newId, Name: name}, nil
		}

		tenant, err := service.Get(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if tenant.Name != name {
			t.Errorf("expected name %v, got %v", name, tenant.Name)
		}

		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestTenantService_Update(t *testing.T) {
	t.Run("should return error when id is empty", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Update(nil, "", "Test Tenant 123", nil, nil)
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Update(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c", "Test Tenant 123", nil, nil)
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not registered", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		repo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		_, err := service.Update(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", "Test Tenant 123", nil, nil)
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should delete tenant successfully", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		repo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse("df5872a7-a907-4e5a-bc5c-23a79e595baa")
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		name := "Test Tenant 123"
		logo := "logo.png"
		banner := "banner.png"
		tenant, err := service.Update(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", name, &logo, &banner)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if tenant.Name != name {
			t.Errorf("expected name %v, got %v", name, tenant.Name)
		}

		if *tenant.Logo != logo {
			t.Errorf("expected logo %v, got %v", logo, tenant.Logo)
		}

		if *tenant.Banner != banner {
			t.Errorf("expected banner %v, got %v", banner, tenant.Banner)
		}
	})
}

func TestTenantService_Delete(t *testing.T) {
	t.Run("should return error when id is empty", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)
		err := service.Delete(nil, "")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)
		err := service.Delete(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not registered", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		repo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		err := service.Delete(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should delete tenant successfully", func(t *testing.T) {
		repo := &repository.TenantRepository{}
		service := NewTenantService(repo)

		repo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse("df5872a7-a907-4e5a-bc5c-23a79e595baa")
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		err := service.Delete(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}
