package domain

import (
	"strings"

	"github.com/google/uuid"
)

type Storage struct {
	ID       uuid.UUID
	Bucket   string
	Filename string
}

func NewStorage(bucket, filename string) (*Storage, error) {
	if strings.TrimSpace(bucket) == "" {
		return nil, ErrBucketRequired
	}

	if strings.TrimSpace(filename) == "" {
		return nil, ErrFilenameRequired
	}

	return &Storage{
		ID:       uuid.New(),
		Bucket:   bucket,
		Filename: filename,
	}, nil
}
