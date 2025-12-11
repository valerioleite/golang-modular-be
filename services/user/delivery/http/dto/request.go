package dto

type CreateUserRequest struct {
	Sub       string  `json:"sub"`
	Email     string  `json:"email"`
	Username  *string `json:"username"`
	FirstName *string `json:"first_name"`
	LastName  *string `json:"last_name"`
}
