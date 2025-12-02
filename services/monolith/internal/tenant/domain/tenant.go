package domain

import (
	"strings"

	"github.com/google/uuid"
)

type ImageType int8

const (
	ImageTypeUnknown ImageType = iota
	ImageTypeLogo
	ImageTypeBanner
)

func (it ImageType) IsValid() bool {
	return it == ImageTypeBanner || it == ImageTypeLogo
}

func ImageTypeFromString(s string) ImageType {
	switch s {
	case "logo":
		return ImageTypeLogo
	case "banner":
		return ImageTypeBanner
	default:
		return ImageTypeUnknown
	}
}

type Tenant struct {
	ID     uuid.UUID
	Name   string
	Logo   *string
	Banner *string
}

func NewTenant(name string) (*Tenant, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrNameRequired
	}

	return &Tenant{
		ID:   uuid.New(),
		Name: name,
	}, nil
}

func (t *Tenant) Update(name string) error {
	if strings.TrimSpace(name) == "" {
		return ErrNameRequired
	}

	t.Name = name

	return nil
}
