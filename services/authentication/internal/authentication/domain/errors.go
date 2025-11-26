package domain

import domainLib "libraries/domain"

var (
	ErrInvalidToken      = domainLib.NewInvalidFieldsError("invalid token")
	ErrTokenExpired      = domainLib.NewInvalidFieldsError("token expired")
	ErrInvalidProvider   = domainLib.NewInvalidFieldsError("invalid OIDC provider configuration")
	ErrAuthenticationFailed = domainLib.NewInvalidFieldsError("authentication failed")
	ErrMissingState      = domainLib.NewInvalidFieldsError("missing state parameter")
	ErrInvalidState      = domainLib.NewInvalidFieldsError("invalid state parameter")
)

