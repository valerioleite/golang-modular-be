package service

import (
	"context"
	"mime/multipart"
	"path/filepath"
	"services/tenant/internal/tenant/domain"
	"services/tenant/internal/tenant/repository"
	"strings"

	"github.com/google/uuid"
)

type TenantService struct {
	tenantRepo  repository.TenantRepository
	storageRepo repository.StorageRepository
}

func NewTenantService(repo repository.TenantRepository, storageRepo repository.StorageRepository) *TenantService {
	return &TenantService{
		tenantRepo:  repo,
		storageRepo: storageRepo,
	}
}

func (s *TenantService) List(ctx context.Context) ([]*domain.Tenant, error) {
	return s.tenantRepo.GetAll(ctx)
}

func (s *TenantService) Get(ctx context.Context, id string) (*domain.Tenant, error) {
	_, err := uuid.Parse(id)
	if err != nil {
		return nil, domain.ErrInvalidID
	}

	tenant, err := s.tenantRepo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if tenant == nil {
		return nil, domain.ErrTenantNotFound
	}

	return tenant, nil
}

func (s *TenantService) Create(ctx context.Context, name string) (*domain.Tenant, error) {
	tenant, err := domain.NewTenant(name)
	if err != nil {
		return nil, err
	}

	if err := s.tenantRepo.Create(ctx, tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *TenantService) Update(ctx context.Context, id, name string) (*domain.Tenant, error) {
	tenant, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	err = tenant.Update(name)
	if err != nil {
		return nil, err
	}

	if err := s.tenantRepo.Update(ctx, tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *TenantService) UpdateImage(ctx context.Context, id string, imageType domain.ImageType, file multipart.File, header *multipart.FileHeader) (*domain.Tenant, error) {
	if file == nil {
		return nil, domain.ErrImageRequired
	}

	if !imageType.IsValid() {
		return nil, domain.ErrInvalidImageType
	}

	if err := s.validateFileType(imageType, header); err != nil {
		return nil, err
	}

	tenant, err := s.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	filename := s.generateFilename(imageType, ext)
	storage, err := domain.NewStorage(id, filename, file)
	if err != nil {
		return nil, err
	}

	path, err := s.storageRepo.Upload(ctx, storage)
	if err != nil {
		return nil, err
	}

	switch imageType {
	case domain.ImageTypeLogo:
		tenant.Logo = path
	case domain.ImageTypeBanner:
		tenant.Banner = path
	default:
		return nil, domain.ErrInvalidImageType
	}

	if err := s.tenantRepo.Update(ctx, tenant); err != nil {
		return nil, err
	}

	return tenant, nil
}

func (s *TenantService) Delete(ctx context.Context, id string) error {
	_, err := s.Get(ctx, id)
	if err != nil {
		return err
	}

	return s.tenantRepo.Delete(ctx, id)
}

func (s *TenantService) validateFileType(imageType domain.ImageType, header *multipart.FileHeader) error {
	if header == nil {
		return domain.ErrInvalidFileType
	}

	ext := strings.ToLower(filepath.Ext(header.Filename))
	contentType := header.Header.Get("Content-Type")

	switch imageType {
	case domain.ImageTypeLogo:
		err := s.isValidLogoFile(ext, contentType)
		if err != nil {
			return err
		}
	case domain.ImageTypeBanner:
		err := s.isValidBannerFile(ext, contentType)
		if err != nil {
			return err
		}
	default:
		return domain.ErrInvalidImageType
	}

	return nil
}

func (s *TenantService) isValidLogoFile(ext, contentType string) error {
	validExtensions := []string{".png", ".jpg", ".jpeg"}
	validContentTypes := []string{"image/png", "image/jpeg", "image/jpg"}

	if !s.isValidExtension(ext, validExtensions) {
		return domain.NewInvalidFileTypeError(validExtensions)
	}

	if contentType != "" && !s.isValidContentType(contentType, validContentTypes) {
		return domain.NewInvalidFileTypeError(validExtensions)
	}

	return nil
}

func (s *TenantService) isValidBannerFile(ext, contentType string) error {
	validExtensions := []string{".jpg", ".jpeg"}
	validContentTypes := []string{"image/jpeg", "image/jpg"}

	if !s.isValidExtension(ext, validExtensions) {
		return domain.NewInvalidFileTypeError(validExtensions)
	}

	if contentType != "" && !s.isValidContentType(contentType, validContentTypes) {
		return domain.NewInvalidFileTypeError(validExtensions)
	}

	return nil
}

func (s *TenantService) isValidExtension(ext string, validExts []string) bool {
	for _, validExt := range validExts {
		if ext == validExt {
			return true
		}
	}

	return false
}

func (s *TenantService) isValidContentType(contentType string, validContentTypes []string) bool {
	for _, validType := range validContentTypes {
		if contentType == validType {
			return true
		}
	}

	return false
}

func (s *TenantService) generateFilename(imageType domain.ImageType, ext string) string {
	var prefix string
	switch imageType {
	case domain.ImageTypeLogo:
		prefix = "logo"
	case domain.ImageTypeBanner:
		prefix = "banner"
	default:
		prefix = "image"
	}

	return prefix + ext
}
