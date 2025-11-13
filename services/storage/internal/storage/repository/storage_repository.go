package repository

import (
	"context"
	"io"
	"services/storage/internal/storage/domain"
)

type StorageRepository interface {
	Init() error
	Upload(ctx context.Context, storage *domain.Storage, file io.Reader) error
}
