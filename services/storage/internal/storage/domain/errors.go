package domain

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidFields = errors.New("invalid fields")
)

type DomainError struct {
	BaseError error
	Message   string
}

func (e *DomainError) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.BaseError.Error()
}

func (e *DomainError) HTTPStatus() int {
	if errors.Is(e.BaseError, ErrNotFound) {
		return 404
	}
	if errors.Is(e.BaseError, ErrInvalidFields) {
		return 400
	}
	return 500
}

func NewNotFoundError(message string) *DomainError {
	return &DomainError{
		BaseError: ErrNotFound,
		Message:   message,
	}
}

func NewInvalidFieldsError(message string) *DomainError {
	return &DomainError{
		BaseError: ErrInvalidFields,
		Message:   message,
	}
}

var (
	ErrBucketRequired   = NewInvalidFieldsError("bucket is required")
	ErrFilenameRequired = NewInvalidFieldsError("filename is required")
	ErrInvalidFile      = NewInvalidFieldsError("invalid file")
)
