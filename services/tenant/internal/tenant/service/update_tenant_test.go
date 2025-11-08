package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"services/tenant/internal/tenant/domain"
	"testing"
)

func TestTenantService_Update(t *testing.T) {
	t.Run("should return error when id is empty", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Update(nil, "", "Test Tenant 123", nil, nil)
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)
		_, err := service.Update(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c", "Test Tenant 123", nil, nil)
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not registered", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)

		repo.getByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		_, err := service.Update(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", "Test Tenant 123", nil, nil)
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should delete tenant successfully", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)

		repo.getByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
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
