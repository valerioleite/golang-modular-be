package domain

import (
	"mime/multipart"
	"strings"
)

type Storage struct {
	Bucket   string
	Filename string
	File     multipart.File
	Path     string
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
		Bucket:   bucket,
		Filename: filename,
		File:     file,
		Path:     bucket + "/" + filename,
	}, nil
}

func DownloadStorage(bucket, filename string) (*Storage, error) {
	if strings.TrimSpace(bucket) == "" {
		return nil, ErrBucketRequired
	}

	if strings.TrimSpace(filename) == "" {
		return nil, ErrFilenameRequired
	}

	return &Storage{
		Bucket:   bucket,
		Filename: filename,
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
