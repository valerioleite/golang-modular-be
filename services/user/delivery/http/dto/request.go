package dto

type CreateUserRequest struct {
	Sub      string  `json:"sub"`
	Email    string  `json:"email"`
	Name     string  `json:"name"`
	Username *string `json:"username"`
}
