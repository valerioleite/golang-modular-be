package domain

import (
	"time"
)

type Token struct {
	AccessToken  string
	RefreshToken string
	IDToken      string
	ExpiresIn    int64
	TokenType    string
}

type UserInfo struct {
	Subject           string
	Email             string
	EmailVerified     bool
	Name              string
	GivenName         string
	FamilyName        string
	PreferredUsername string
	Picture           string
}

type AuthState struct {
	State       string
	RedirectURI string
	CreatedAt   time.Time
}

func NewAuthState(state, redirectURI string) *AuthState {
	return &AuthState{
		State:       state,
		RedirectURI: redirectURI,
		CreatedAt:   time.Now(),
	}
}

func (a *AuthState) IsExpired(maxAge time.Duration) bool {
	return time.Since(a.CreatedAt) > maxAge
}
