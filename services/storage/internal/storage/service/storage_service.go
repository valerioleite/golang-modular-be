package service

import (
	"context"
	"mime/multipart"
	"services/storage/internal/storage/domain"
	"services/storage/internal/storage/repository"
)

type StorageService struct {
	repo repository.StorageRepository
}

func NewStorageService(repo repository.StorageRepository) *StorageService {
	return &StorageService{repo: repo}
}

func (s *StorageService) Init() error {
	return s.repo.Init()
}

func (s *StorageService) Upload(ctx context.Context, bucket, filename string, file multipart.File, header *multipart.FileHeader) (*domain.Storage, error) {
	storage, err := domain.NewStorage(bucket, filename, file, header)
	if err != nil {
		return nil, err
	}

	if err := s.repo.Upload(ctx, storage, file); err != nil {
		return nil, err
	}

	return storage, nil
}
