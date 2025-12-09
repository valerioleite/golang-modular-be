package domain

import (
	"strings"
	"time"

	"github.com/google/uuid"
)

type User struct {
	ID        uuid.UUID
	CreatedBy string
	CreatedAt time.Time
	UpdatedBy string
	UpdatedAt time.Time
	Sub       string
	Email     string
	Name      string
	Username  *string
}

func NewUser(createdBy, sub, email, name string, username *string) (*User, error) {
	if strings.TrimSpace(sub) == "" {
		return nil, ErrSubRequired
	}

	if strings.TrimSpace(email) == "" {
		return nil, ErrEmailRequired
	}

	if strings.TrimSpace(name) == "" {
		return nil, ErrNameRequired
	}

	now := time.Now()
	return &User{
		ID:        uuid.New(),
		CreatedBy: createdBy,
		CreatedAt: now,
		UpdatedBy: createdBy,
		UpdatedAt: now,
		Sub:       sub,
		Email:     email,
		Name:      name,
		Username:  username,
	}, nil
}
