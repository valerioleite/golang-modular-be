package domain

import (
	"mime/multipart"
	"strings"
)

type Storage struct {
	Bucket   string
	Filename string
	File     multipart.File
}

func NewStorage(bucket, filename string, file multipart.File) (*Storage, error) {
	if strings.TrimSpace(bucket) == "" {
		return nil, ErrBucketRequired
	}

	if strings.TrimSpace(filename) == "" {
		return nil, ErrFilenameRequired
	}

	return &Storage{
		Bucket:   bucket,
		Filename: filename,
		File:     file,
	}, nil
}
