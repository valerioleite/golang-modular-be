package domain

import (
	"fmt"
	domainLib "libraries/domain"
	"strings"
)

var (
	ErrTenantNotFound   = domainLib.NewNotFoundError("tenant not found")
	ErrInvalidID        = domainLib.NewInvalidFieldsError("invalid id")
	ErrInvalidImageType = domainLib.NewInvalidFieldsError("invalid image type")
	ErrInvalidFileType  = domainLib.NewInvalidFieldsError("invalid file type")
	ErrNameRequired     = domainLib.NewInvalidFieldsError("name is required")
	ErrBucketRequired   = domainLib.NewInvalidFieldsError("bucket is required")
	ErrImageRequired    = domainLib.NewInvalidFieldsError("image is required")
	ErrFilenameRequired = domainLib.NewInvalidFieldsError("filename is required")
)

func NewInvalidFileTypeError(validExtensions []string) error {
	extensionsStr := strings.Join(validExtensions, ", ")
	message := fmt.Sprintf("invalid file type. Accepted extensions: %s", extensionsStr)
	return domainLib.NewInvalidFieldsError(message)
}
