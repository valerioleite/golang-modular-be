package http

import (
	"libraries/http/middleware"
	"net/http"
	"services/monolith/internal/authentication/delivery/http/handlers"
	"services/monolith/internal/authentication/service"
)

type Router struct {
	loginHandler        *handlers.LoginHandler
	callbackGetHandler  *handlers.CallbackGetHandler
	callbackPostHandler *handlers.CallbackPostHandler
	refreshTokenHandler *handlers.RefreshTokenHandler
	verifyTokenHandler  *handlers.VerifyTokenHandler
	userInfoHandler     *handlers.UserInfoHandler
}

func NewRouter(service *service.AuthenticationService) *Router {
	return &Router{
		loginHandler:        handlers.NewLoginHandler(service),
		callbackGetHandler:  handlers.NewCallbackGetHandler(service),
		callbackPostHandler: handlers.NewCallbackPostHandler(service),
		refreshTokenHandler: handlers.NewRefreshTokenHandler(service),
		verifyTokenHandler:  handlers.NewVerifyTokenHandler(service),
		userInfoHandler:     handlers.NewUserInfoHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /auth/login", r.loginHandler.Handle)
	mux.HandleFunc("GET /auth/callback", r.callbackGetHandler.Handle)
	mux.HandleFunc("POST /auth/callback", r.callbackPostHandler.Handle)
	mux.HandleFunc("POST /auth/refresh", r.refreshTokenHandler.Handle)
	mux.HandleFunc("POST /auth/verify", r.verifyTokenHandler.Handle)
	mux.Handle("GET /auth/userinfo", middleware.HandleWithValidateToken(r.userInfoHandler.Handle))

	return mux
}
