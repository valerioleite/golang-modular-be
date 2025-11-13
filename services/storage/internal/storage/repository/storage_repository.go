package repository

import (
	"context"
	"io"
	"services/storage/internal/storage/domain"
)

type StorageRepository interface {
	Init() error
	ExistsBucket(ctx context.Context, bucket string) bool
	CreateBucket(ctx context.Context, bucket string) error
	Upload(ctx context.Context, storage *domain.Storage, file io.Reader) error
}
