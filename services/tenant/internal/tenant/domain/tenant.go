package domain

import (
	"strings"

	"github.com/google/uuid"
)

type Tenant struct {
	ID     uuid.UUID
	Name   string
	Logo   *string
	Banner *string
}

func NewTenant(name string, logo, banner *string) (*Tenant, error) {
	if strings.TrimSpace(name) == "" {
		return nil, ErrNameRequired
	}

	return &Tenant{
		ID:     uuid.New(),
		Name:   name,
		Logo:   logo,
		Banner: banner,
	}, nil
}

func (t *Tenant) Update(name string, logo, banner *string) error {
	if strings.TrimSpace(name) == "" {
		return ErrNameRequired
	}

	t.Name = name

	if logo != nil {
		t.Logo = logo
	}

	if banner != nil {
		t.Banner = banner
	}

	return nil
}
