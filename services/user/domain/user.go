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
	Username  *string
	FirstName *string
	LastName  *string
}

func NewUser(createdBy, sub, email string, username, firstName, lastName *string) (*User, error) {
	if strings.TrimSpace(sub) == "" {
		return nil, ErrSubRequired
	}

	if strings.TrimSpace(email) == "" {
		return nil, ErrEmailRequired
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
		Username:  username,
		FirstName: firstName,
		LastName:  lastName,
	}, nil
}
