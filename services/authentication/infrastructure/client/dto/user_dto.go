package dto

import "time"

type CreateUserRequest struct {
	Sub       string  `json:"sub"`
	Email     string  `json:"email"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
	Username  *string `json:"username"`
}

type UserResponse struct {
	ID        string    `json:"id"`
	CreatedBy string    `json:"created_by"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedBy string    `json:"updated_by"`
	UpdatedAt time.Time `json:"updated_at"`
	Sub       string    `json:"sub"`
	Email     string    `json:"email"`
	Username  *string   `json:"username,omitempty"`
	FirstName *string   `json:"first_name,omitempty"`
	LastName  *string   `json:"last_name,omitempty"`
}
