package repository

import (
	"context"
	"services/tenant/domain"
	"services/tenant/infrastructure/client"
	"services/tenant/repository"
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

	return &resp.Path, nil
}
