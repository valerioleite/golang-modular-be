package oidc

import (
	"context"
	"fmt"
	"os"
	"services/authentication/domain"
	"services/authentication/repository"
	"time"

	"github.com/coreos/go-oidc/v3/oidc"
	"golang.org/x/oauth2"
)

var (
	issuerUrlKey    = "OIDC_ISSUER_URL"
	clientIdKey     = "OIDC_CLIENT_ID"
	clientSecretKey = "OIDC_CLIENT_SECRET"
	redirectUriKey  = "OIDC_REDIRECT_URI"
)

type AuthenticationRepositoryOIDC struct {
	provider     *oidc.Provider
	verifier     *oidc.IDTokenVerifier
	oauth2Config *oauth2.Config
	issuerURL    string
	clientID     string
	clientSecret string
	redirectURI  string
}

func NewOIDCRepository() repository.AuthenticationRepository {
	return &AuthenticationRepositoryOIDC{
		issuerURL:    os.Getenv(issuerUrlKey),
		clientID:     os.Getenv(clientIdKey),
		clientSecret: os.Getenv(clientSecretKey),
		redirectURI:  os.Getenv(redirectUriKey),
	}
}

func (r *AuthenticationRepositoryOIDC) Init(ctx context.Context) error {
	if r.issuerURL == "" {
		return domain.ErrInvalidProvider
	}

	provider, err := oidc.NewProvider(ctx, r.issuerURL)
	if err != nil {
		return fmt.Errorf("failed to create OIDC provider: %w", err)
	}

	r.provider = provider
	r.verifier = provider.Verifier(&oidc.Config{
		ClientID: r.clientID,
	})

	r.oauth2Config = &oauth2.Config{
		ClientID:     r.clientID,
		ClientSecret: r.clientSecret,
		Endpoint:     provider.Endpoint(),
		RedirectURL:  r.redirectURI,
		Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "offline_access"},
	}

	return nil
}

func (r *AuthenticationRepositoryOIDC) GetAuthorizationURL(ctx context.Context, state string) (string, error) {
	if r.oauth2Config == nil {
		return "", domain.ErrInvalidProvider
	}

	config := *r.oauth2Config
	authURL := config.AuthCodeURL(state, oauth2.AccessTypeOffline)

	return authURL, nil
}

func (r *AuthenticationRepositoryOIDC) ExchangeCode(ctx context.Context, code string) (*domain.Token, error) {
	if r.oauth2Config == nil {
		return nil, domain.ErrInvalidProvider
	}

	config := *r.oauth2Config
	oauth2Token, err := config.Exchange(ctx, code)
	if err != nil {
		return nil, fmt.Errorf("failed to exchange code: %w", err)
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		return nil, domain.ErrInvalidToken
	}

	token := &domain.Token{
		AccessToken:  oauth2Token.AccessToken,
		RefreshToken: oauth2Token.RefreshToken,
		IDToken:      rawIDToken,
		TokenType:    oauth2Token.TokenType,
	}

	if expiry := oauth2Token.Expiry; !expiry.IsZero() {
		token.ExpiresIn = int64(time.Until(expiry).Seconds())
	}

	return token, nil
}

func (r *AuthenticationRepositoryOIDC) RefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	if r.oauth2Config == nil {
		return nil, domain.ErrInvalidProvider
	}

	tokenSource := r.oauth2Config.TokenSource(ctx, &oauth2.Token{
		RefreshToken: refreshToken,
	})

	oauth2Token, err := tokenSource.Token()
	if err != nil {
		return nil, fmt.Errorf("failed to refresh token: %w", err)
	}

	rawIDToken, ok := oauth2Token.Extra("id_token").(string)
	if !ok {
		rawIDToken = ""
	}

	token := &domain.Token{
		AccessToken:  oauth2Token.AccessToken,
		RefreshToken: oauth2Token.RefreshToken,
		IDToken:      rawIDToken,
		TokenType:    oauth2Token.TokenType,
	}

	if expiry := oauth2Token.Expiry; !expiry.IsZero() {
		token.ExpiresIn = int64(time.Until(expiry).Seconds())
	}

	return token, nil
}

func (r *AuthenticationRepositoryOIDC) VerifyToken(ctx context.Context, token string) (*domain.UserInfo, error) {
	if r.verifier == nil {
		return nil, domain.ErrInvalidProvider
	}

	idToken, err := r.verifier.Verify(ctx, token)
	if err != nil {
		return nil, fmt.Errorf("failed to verify token: %w", err)
	}

	var claims Claims
	if err := idToken.Claims(&claims); err != nil {
		return nil, fmt.Errorf("failed to parse claims: %w", err)
	}

	return &domain.UserInfo{
		Subject:           claims.Subject,
		Email:             claims.Email,
		EmailVerified:     claims.EmailVerified,
		Name:              claims.Name,
		GivenName:         claims.GivenName,
		FamilyName:        claims.FamilyName,
		PreferredUsername: claims.PreferredUsername,
		Picture:           claims.Picture,
	}, nil
}
