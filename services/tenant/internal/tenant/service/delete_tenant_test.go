package service

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"services/tenant/internal/tenant/domain"
	"testing"
)

func TestTenantService_Delete(t *testing.T) {
	t.Run("should return error when id is empty", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)
		err := service.Delete(nil, "")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		repo := &mockTenantRepository{}
		service := NewTenantService(repo)
		err := service.Delete(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c")
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

		err := service.Delete(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
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

		err := service.Delete(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})

}
