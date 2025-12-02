package domain

import domainLib "libraries/domain"

var (
	ErrInvalidToken       = domainLib.NewInvalidFieldsError("invalid token")
	ErrInvalidProvider    = domainLib.NewInvalidFieldsError("invalid OIDC provider configuration")
	ErrMissingCodeOrState = domainLib.NewInvalidFieldsError("missing code or state parameter")
	ErrInvalidState       = domainLib.NewInvalidFieldsError("invalid state parameter")
)
