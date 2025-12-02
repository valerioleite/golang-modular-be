package repository

import (
	"context"
	"io"
	"services/monolith/internal/storage/domain"
)

type StorageRepository interface {
	Init() error
	ExistsBucket(ctx context.Context, bucket string) bool
	CreateBucket(ctx context.Context, bucket string) error
	Upload(ctx context.Context, storage *domain.Storage, file io.Reader) error
	Download(ctx context.Context, storage *domain.Storage) (io.ReadCloser, error)
}
