package service

import (
	"context"
	"io"
	"mime/multipart"
	"services/storage/domain"
	"services/storage/repository"
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

	if !s.repo.ExistsBucket(ctx, bucket) {
		if err := s.repo.CreateBucket(ctx, bucket); err != nil {
			return nil, err
		}
	}

	if err := s.repo.Upload(ctx, storage, file); err != nil {
		return nil, err
	}

	return storage, nil
}

func (s *StorageService) Download(ctx context.Context, bucket, filename string) (io.ReadCloser, error) {
	storage, err := domain.DownloadStorage(bucket, filename)
	if err != nil {
		return nil, err
	}

	return s.repo.Download(ctx, storage)
}
