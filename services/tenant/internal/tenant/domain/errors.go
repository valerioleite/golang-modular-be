package domain

import domainLib "libraries/domain"

var (
	ErrTenantNotFound = domainLib.NewNotFoundError("tenant not found")
	ErrInvalidID      = domainLib.NewInvalidFieldsError("invalid id")
	ErrNameRequired   = domainLib.NewInvalidFieldsError("name is required")
)
