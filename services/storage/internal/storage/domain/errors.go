package domain

import domainLib "libraries/domain"

var (
	ErrBucketRequired   = domainLib.NewInvalidFieldsError("bucket is required")
	ErrFilenameRequired = domainLib.NewInvalidFieldsError("filename is required")
	ErrInvalidFile      = domainLib.NewInvalidFieldsError("invalid file")
)
