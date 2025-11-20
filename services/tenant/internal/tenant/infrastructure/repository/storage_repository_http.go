package repository

import (
	"context"
	"services/tenant/internal/tenant/domain"
	"services/tenant/internal/tenant/infrastructure/client"
	"services/tenant/internal/tenant/repository"
)

type StorageRepositoryHttp struct {
	client *client.StorageClient
}

func NewStorageRepositoryHttp(client *client.StorageClient) repository.StorageRepository {
	return &StorageRepositoryHttp{client: client}
}

func (s StorageRepositoryHttp) Upload(ctx context.Context, storage *domain.Storage) (*string, error) {
	resp, err := s.client.UploadFile(ctx, storage.Bucket, storage.Filename, storage.File)
	if err != nil {
		return nil, err
	}

	return &resp.Filename, nil
}
