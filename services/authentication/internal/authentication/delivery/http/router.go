package http

import (
	"net/http"
	"services/authentication/internal/authentication/delivery/http/handlers"
	"services/authentication/internal/authentication/service"
)

type Router struct {
	loginHandler        *handlers.LoginHandler
	callbackGetHandler  *handlers.CallbackGetHandler
	refreshTokenHandler *handlers.RefreshTokenHandler
	verifyTokenHandler  *handlers.VerifyTokenHandler
	userInfoHandler     *handlers.UserInfoHandler
}

func NewRouter(service *service.AuthenticationService, frontendURL string) *Router {
	return &Router{
		loginHandler:        handlers.NewLoginHandler(service),
		callbackGetHandler:  handlers.NewCallbackGetHandler(service, frontendURL),
		refreshTokenHandler: handlers.NewRefreshTokenHandler(service),
		verifyTokenHandler:  handlers.NewVerifyTokenHandler(service),
		userInfoHandler:     handlers.NewUserInfoHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", r.loginHandler.Handle)
	mux.HandleFunc("GET /auth/callback", r.callbackGetHandler.Handle)
	mux.HandleFunc("POST /auth/refresh", r.refreshTokenHandler.Handle)
	mux.HandleFunc("POST /auth/verify", r.verifyTokenHandler.Handle)
	mux.HandleFunc("GET /auth/userinfo", r.userInfoHandler.Handle)

	return mux
}
