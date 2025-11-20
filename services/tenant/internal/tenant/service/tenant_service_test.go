package service

import (
	"bytes"
	"context"
	"errors"
	"mime/multipart"
	"services/tenant/internal/tenant/domain"
	"services/tenant/internal/tenant/infrastructure/repository"
	"testing"

	"github.com/google/uuid"
)

// mockFile implementa multipart.File para testes
type mockFile struct {
	*bytes.Reader
	closed bool
}

func (m *mockFile) Close() error {
	m.closed = true
	return nil
}

func newMockFile(content []byte) *mockFile {
	return &mockFile{
		Reader: bytes.NewReader(content),
		closed: false,
	}
}

func TestTenantService_Create(t *testing.T) {
	t.Run("should return error when name is empty", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		_, err := service.Create(nil, "")
		if !errors.Is(err, domain.ErrNameRequired) {
			t.Errorf("expected error %v, got %v", domain.ErrNameRequired, err)
		}
	})

	t.Run("should create tenant successfully", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		name := "Tenant X"
		tenant, err := service.Create(nil, name)
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
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetAllFunc = func(ctx context.Context) ([]*domain.Tenant, error) {
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
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetAllFunc = func(ctx context.Context) ([]*domain.Tenant, error) {
			tenant1, _ := domain.NewTenant("Tenant 1")
			tenant2, _ := domain.NewTenant("Tenant 2")
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
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		_, err := service.Get(nil, "")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		_, err := service.Get(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not registered", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		_, err := service.Get(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should get tenant successfully", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		name := "Test Tenant 123"

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
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
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		_, err := service.Update(nil, "", "Test Tenant 123")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		_, err := service.Update(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c", "Test Tenant 123")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not registered", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		_, err := service.Update(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", "Test Tenant 123")
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should delete tenant successfully", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse("df5872a7-a907-4e5a-bc5c-23a79e595baa")
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		name := "Test Tenant 123"
		tenant, err := service.Update(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", name)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if tenant.Name != name {
			t.Errorf("expected name %v, got %v", name, tenant.Name)
		}
	})
}

func TestTenantService_Delete(t *testing.T) {
	t.Run("should return error when id is empty", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		err := service.Delete(nil, "")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not uuid", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		err := service.Delete(nil, "550a5288-8e65-450c-bd2c-0028d4a1d3c")
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when id is not registered", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		err := service.Delete(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should delete tenant successfully", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse("df5872a7-a907-4e5a-bc5c-23a79e595baa")
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		err := service.Delete(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa")
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}
	})
}

func TestTenantService_UpdateImage(t *testing.T) {
	t.Run("should return error when file is nil", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		header := &multipart.FileHeader{
			Filename: "test.png",
		}

		_, err := service.UpdateImage(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", domain.ImageTypeLogo, nil, header)
		if !errors.Is(err, domain.ErrImageRequired) {
			t.Errorf("expected error %v, got %v", domain.ErrImageRequired, err)
		}
	})

	t.Run("should return error when imageType is invalid", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		file := newMockFile([]byte("test content"))
		header := &multipart.FileHeader{
			Filename: "test.png",
		}

		_, err := service.UpdateImage(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", domain.ImageTypeUnknown, file, header)
		if !errors.Is(err, domain.ErrInvalidImageType) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidImageType, err)
		}
	})

	t.Run("should return error when id is invalid", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		file := newMockFile([]byte("test content"))
		header := &multipart.FileHeader{
			Filename: "test.png",
		}

		_, err := service.UpdateImage(nil, "", domain.ImageTypeLogo, file, header)
		if !errors.Is(err, domain.ErrInvalidID) {
			t.Errorf("expected error %v, got %v", domain.ErrInvalidID, err)
		}
	})

	t.Run("should return error when tenant is not found", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			return nil, nil
		}

		file := newMockFile([]byte("test content"))
		header := &multipart.FileHeader{
			Filename: "test.png",
		}

		_, err := service.UpdateImage(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", domain.ImageTypeLogo, file, header)
		if !errors.Is(err, domain.ErrTenantNotFound) {
			t.Errorf("expected error %v, got %v", domain.ErrTenantNotFound, err)
		}
	})

	t.Run("should return error when logo file extension is invalid", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse("df5872a7-a907-4e5a-bc5c-23a79e595baa")
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		file := newMockFile([]byte("test content"))
		header := &multipart.FileHeader{
			Filename: "test.gif",
		}

		_, err := service.UpdateImage(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", domain.ImageTypeLogo, file, header)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should return error when banner file extension is invalid", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse("df5872a7-a907-4e5a-bc5c-23a79e595baa")
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		file := newMockFile([]byte("test content"))
		header := &multipart.FileHeader{
			Filename: "test.png",
		}

		_, err := service.UpdateImage(nil, "df5872a7-a907-4e5a-bc5c-23a79e595baa", domain.ImageTypeBanner, file, header)
		if err == nil {
			t.Errorf("expected error, got nil")
		}
	})

	t.Run("should update logo successfully", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		tenantID := "df5872a7-a907-4e5a-bc5c-23a79e595baa"

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse(tenantID)
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		storageRepo.UploadFunc = func(ctx context.Context, storage *domain.Storage) (*string, error) {
			path := "test/path/" + storage.Filename
			return &path, nil
		}

		tenantRepo.UpdateFunc = func(ctx context.Context, tenant *domain.Tenant) error {
			return nil
		}

		file := newMockFile([]byte("test content"))
		header := &multipart.FileHeader{
			Filename: "test.png",
		}

		tenant, err := service.UpdateImage(nil, tenantID, domain.ImageTypeLogo, file, header)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if tenant.Logo == nil {
			t.Errorf("expected logo to be set, got nil")
		}

		if tenant.Logo != nil && *tenant.Logo != "test/path/logo.png" {
			t.Errorf("expected logo path 'test/path/logo.png', got %v", *tenant.Logo)
		}
	})

	t.Run("should update banner successfully", func(t *testing.T) {
		tenantRepo := &repository.TenantRepository{}
		storageRepo := &repository.StorageRepositoryMock{}
		service := NewTenantService(tenantRepo, storageRepo)
		tenantID := "df5872a7-a907-4e5a-bc5c-23a79e595baa"

		tenantRepo.GetByIdFunc = func(ctx context.Context, id string) (*domain.Tenant, error) {
			newId, _ := uuid.Parse(tenantID)
			return &domain.Tenant{ID: newId, Name: "Test Tenant"}, nil
		}

		storageRepo.UploadFunc = func(ctx context.Context, storage *domain.Storage) (*string, error) {
			path := "test/path/" + storage.Filename
			return &path, nil
		}

		tenantRepo.UpdateFunc = func(ctx context.Context, tenant *domain.Tenant) error {
			return nil
		}

		file := newMockFile([]byte("test content"))
		header := &multipart.FileHeader{
			Filename: "test.jpg",
		}

		tenant, err := service.UpdateImage(nil, tenantID, domain.ImageTypeBanner, file, header)
		if err != nil {
			t.Errorf("unexpected error: %v", err)
		}

		if tenant.Banner == nil {
			t.Errorf("expected banner to be set, got nil")
		}

		if tenant.Banner != nil && *tenant.Banner != "test/path/banner.jpg" {
			t.Errorf("expected banner path 'test/path/banner.jpg', got %v", *tenant.Banner)
		}
	})
}
