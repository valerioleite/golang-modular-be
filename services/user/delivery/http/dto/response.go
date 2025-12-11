package dto

type UserResponse struct {
	ID        string  `json:"id"`
	CreatedBy string  `json:"created_by"`
	CreatedAt string  `json:"created_at"`
	UpdatedBy string  `json:"updated_by"`
	UpdatedAt string  `json:"updated_at"`
	Sub       string  `json:"sub"`
	Email     string  `json:"email"`
	Username  *string `json:"username,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
} //@name UserResponse
