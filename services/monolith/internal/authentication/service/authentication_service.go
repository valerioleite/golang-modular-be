package service

import (
	"context"
	"services/monolith/internal/authentication/domain"
	"services/monolith/internal/authentication/repository"
	"sync"
	"time"

	"github.com/google/uuid"
)

type AuthenticationService struct {
	oidcRepo repository.AuthenticationRepository
	states   map[string]*domain.AuthState
	mu       sync.RWMutex
}

func NewAuthenticationService(oidcRepo repository.AuthenticationRepository) *AuthenticationService {
	service := &AuthenticationService{
		oidcRepo: oidcRepo,
		states:   make(map[string]*domain.AuthState),
	}

	go service.cleanupExpiredStates()

	return service
}

func (s *AuthenticationService) Init(ctx context.Context) error {
	return s.oidcRepo.Init(ctx)
}

func (s *AuthenticationService) Login(ctx context.Context, redirectURI string) (string, string, error) {
	state, err := generateState()
	if err != nil {
		return "", "", err
	}

	authURL, err := s.oidcRepo.GetAuthorizationURL(ctx, state)
	if err != nil {
		return "", "", err
	}

	s.storeState(state, redirectURI)

	return authURL, state, nil
}

func (s *AuthenticationService) Callback(ctx context.Context, code, state string) (*domain.Token, error) {
	if code == "" || state == "" {
		return nil, domain.ErrMissingCodeOrState
	}

	if !s.validateState(state) {
		return nil, domain.ErrInvalidState
	}

	token, err := s.oidcRepo.ExchangeCode(ctx, code)
	if err != nil {
		return nil, err
	}

	s.removeState(state)

	return token, nil
}

func (s *AuthenticationService) RefreshToken(ctx context.Context, refreshToken string) (*domain.Token, error) {
	return s.oidcRepo.RefreshToken(ctx, refreshToken)
}

func (s *AuthenticationService) VerifyToken(ctx context.Context, token string) (*domain.UserInfo, error) {
	return s.oidcRepo.VerifyToken(ctx, token)
}

func (s *AuthenticationService) GetUserInfo(ctx context.Context, accessToken string) (*domain.UserInfo, error) {
	return s.oidcRepo.GetUserInfo(ctx, accessToken)
}

func (s *AuthenticationService) storeState(state, redirectURI string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.states[state] = domain.NewAuthState(state, redirectURI)
}

func (s *AuthenticationService) validateState(state string) bool {
	s.mu.RLock()
	defer s.mu.RUnlock()

	authState, exists := s.states[state]
	if !exists {
		return false
	}

	if authState.IsExpired(10 * time.Minute) {
		return false
	}

	return true
}

func (s *AuthenticationService) removeState(state string) {
	s.mu.Lock()
	defer s.mu.Unlock()
	delete(s.states, state)
}

func (s *AuthenticationService) cleanupExpiredStates() {
	ticker := time.NewTicker(5 * time.Minute)
	defer ticker.Stop()

	for range ticker.C {
		s.mu.Lock()
		now := time.Now()
		for state, authState := range s.states {
			if authState.IsExpired(10 * time.Minute) {
				delete(s.states, state)
			}
		}
		s.mu.Unlock()
		_ = now
	}
}

func generateState() (string, error) {
	id, err := uuid.NewRandom()
	if err != nil {
		return "", err
	}
	return id.String(), nil
}
