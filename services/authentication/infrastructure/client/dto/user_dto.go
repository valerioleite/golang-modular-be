package dto

type CreateUserRequest struct {
	Sub      string  `json:"sub"`
	Email    string  `json:"email"`
	Name     string  `json:"name"`
	Username *string `json:"username"`
}

type UserResponse struct {
	ID        string  `json:"id"`
	CreatedBy string  `json:"createdBy"`
	CreatedAt string  `json:"createdAt"`
	UpdatedBy string  `json:"updatedBy"`
	UpdatedAt string  `json:"updatedAt"`
	Sub       string  `json:"sub"`
	Email     string  `json:"email"`
	Name      string  `json:"name"`
	Username  *string `json:"username,omitempty"`
}

