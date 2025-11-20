package repository

import (
	"context"
	"services/tenant/internal/tenant/domain"
)

type StorageRepositoryMock struct {
	UploadFunc func(ctx context.Context, storage *domain.Storage) (*string, error)
}

func (m *StorageRepositoryMock) Upload(ctx context.Context, storage *domain.Storage) (*string, error) {
	if m.UploadFunc != nil {
		return m.UploadFunc(ctx, storage)
	}
	path := "test/path/" + storage.Filename
	return &path, nil
}

