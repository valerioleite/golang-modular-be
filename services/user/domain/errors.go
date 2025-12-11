package domain

import (
	domainLib "libraries/domain"
)

var (
	ErrSubRequired       = domainLib.NewInvalidFieldsError("sub is required")
	ErrEmailRequired     = domainLib.NewInvalidFieldsError("email is required")
	ErrInvalidID         = domainLib.NewInvalidFieldsError("invalid id")
	ErrUserNotFound      = domainLib.NewNotFoundError("user not found")
	ErrUserAlreadyExists = domainLib.NewConflictError("user already exists")
)
