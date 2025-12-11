package dto

type LoginRequest struct {
	RedirectURI string `json:"redirect_uri"`
}

type LoginResponse struct {
	AuthURL *string `json:"auth_url"`
} //@name LoginResponse

type CallbackRequest struct {
	Code        string `json:"code"`
	State       string `json:"state"`
	RedirectURI string `json:"redirect_uri"`
} //@name CallbackRequest

type TokenResponse struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	IDToken      string `json:"id_token"`
	ExpiresIn    int64  `json:"expires_in"`
	TokenType    string `json:"token_type"`
} //@name TokenResponse

type RefreshTokenRequest struct {
	RefreshToken string `json:"refresh_token"`
} //@name RefreshTokenRequest

type UserInfoResponse struct {
	Sub       string  `json:"sub"`
	Email     string  `json:"email"`
	Username  *string `json:"username,omitempty"`
	FirstName *string `json:"first_name,omitempty"`
	LastName  *string `json:"last_name,omitempty"`
} //@name UserInfoResponse
