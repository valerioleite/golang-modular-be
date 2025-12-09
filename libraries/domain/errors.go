package domain

import (
	"errors"
)

var (
	ErrNotFound      = errors.New("not found")
	ErrInvalidFields = errors.New("invalid fields")
	ErrConflict      = errors.New("conflict")
)

type Error struct {
	BaseError error
	Message   string
}

func (e *Error) Error() string {
	if e.Message != "" {
		return e.Message
	}
	return e.BaseError.Error()
}

func (e *Error) HTTPStatus() int {
	if errors.Is(e.BaseError, ErrNotFound) {
		return 404
	}
	if errors.Is(e.BaseError, ErrInvalidFields) {
		return 400
	}
	if errors.Is(e.BaseError, ErrConflict) {
		return 409
	}
	return 500
}

func NewNotFoundError(message string) *Error {
	return &Error{
		BaseError: ErrNotFound,
		Message:   message,
	}
}

func NewInvalidFieldsError(message string) *Error {
	return &Error{
		BaseError: ErrInvalidFields,
		Message:   message,
	}
}

func NewConflictError(message string) *Error {
	return &Error{
		BaseError: ErrConflict,
		Message:   message,
	}
}
