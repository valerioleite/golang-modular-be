package domain

import (
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

func UserInfoToUser(userInfo *UserInfo) *User {
	now := time.Now()
	var username *string
	if userInfo.PreferredUsername != "" {
		username = &userInfo.PreferredUsername
	}

	return &User{
		ID:        uuid.New(),
		CreatedBy: userInfo.Subject,
		CreatedAt: now,
		UpdatedBy: userInfo.Subject,
		UpdatedAt: now,
		Sub:       userInfo.Subject,
		Email:     userInfo.Email,
		FirstName: &userInfo.GivenName,
		LastName:  &userInfo.FamilyName,
		Username:  username,
	}
}
