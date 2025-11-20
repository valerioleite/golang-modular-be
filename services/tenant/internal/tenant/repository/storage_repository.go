package repository

import (
	"context"
	"services/tenant/internal/tenant/domain"
)

type StorageRepository interface {
	Upload(ctx context.Context, storage *domain.Storage) (*string, error)
}
