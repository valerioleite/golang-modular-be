package domain

import (
	"mime/multipart"
	"strings"

	"github.com/google/uuid"
)

type Storage struct {
	ID       uuid.UUID
	Bucket   string
	Filename string
	File     multipart.File
}

func NewStorage(bucket, filename string, file multipart.File, header *multipart.FileHeader) (*Storage, error) {
	filename = replaceFilenameToHeader(filename, header)

	if file == nil {
		return nil, ErrInvalidFile
	}

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
		File:     file,
	}, nil
}

func replaceFilenameToHeader(filename string, header *multipart.FileHeader) string {
	if strings.TrimSpace(filename) != "" {
		return filename
	}

	if header == nil {
		return ""
	}

	return header.Filename
}
