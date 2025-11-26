package repository

import (
	"context"
	"services/authentication/internal/authentication/domain"
)

type AuthenticationRepository interface {
	Init(ctx context.Context) error
	GetAuthorizationURL(ctx context.Context, state string) (string, error)
	ExchangeCode(ctx context.Context, code string) (*domain.Token, error)
	RefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error)
	VerifyToken(ctx context.Context, token string) (*domain.UserInfo, error)
	GetUserInfo(ctx context.Context, accessToken string) (*domain.UserInfo, error)
}
