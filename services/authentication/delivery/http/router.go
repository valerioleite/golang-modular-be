package http

import (
	"libraries/http/health"
	"libraries/http/middleware"
	"libraries/http/swagger"
	"net/http"
	"services/authentication/delivery/http/handlers"
	_ "services/authentication/docs"
	"services/authentication/service"
)

type Router struct {
	loginHandler        *handlers.AuthorizeHandler
	callbackGetHandler  *handlers.CallbackGetHandler
	callbackPostHandler *handlers.CallbackPostHandler
	refreshTokenHandler *handlers.RefreshTokenHandler
	verifyTokenHandler  *handlers.VerifyTokenHandler
	userInfoHandler     *handlers.UserInfoHandler
}

func NewRouter(service *service.AuthenticationService) *Router {
	return &Router{
		loginHandler:        handlers.NewAuthorizeHandler(service),
		callbackGetHandler:  handlers.NewCallbackGetHandler(service),
		callbackPostHandler: handlers.NewCallbackPostHandler(service),
		refreshTokenHandler: handlers.NewRefreshTokenHandler(service),
		verifyTokenHandler:  handlers.NewVerifyTokenHandler(service),
		userInfoHandler:     handlers.NewUserInfoHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	loggingMiddleware := middleware.WithLogging("authentication")

	mux.HandleFunc("GET /api/authentication/swagger/", swagger.Handler("authentication"))

	r.setupActuatorResources(mux, loggingMiddleware)
	r.setupAuthenticationResources(mux, loggingMiddleware)

	return mux
}

// healthCheck godoc
// @Summary Health check
// @Description Returns the health status of the authentication module.
// @Tags Actuator
// @Produce json
// @Success 200 {object} dto.HealthResponse
// @Router /v1/actuator/health [get]
func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	health.Handler("authentication")(w, req)
}

func (r *Router) setupActuatorResources(mux *http.ServeMux, loggingMiddleware func(http.Handler) http.Handler) {
	mux.Handle("GET /api/authentication/v1/actuator/health", loggingMiddleware(http.HandlerFunc(r.healthCheck)))
}

func (r *Router) setupAuthenticationResources(mux *http.ServeMux, loggingMiddleware func(http.Handler) http.Handler) {
	mux.Handle("GET /api/authentication/v1/authentication/authorize", loggingMiddleware(http.HandlerFunc(r.loginHandler.Handle)))
	mux.Handle("GET /api/authentication/v1/authentication/callback", loggingMiddleware(http.HandlerFunc(r.callbackGetHandler.Handle)))
	mux.Handle("POST /api/authentication/v1/authentication/callback", loggingMiddleware(http.HandlerFunc(r.callbackPostHandler.Handle)))
	mux.Handle("POST /api/authentication/v1/authentication/refresh", loggingMiddleware(http.HandlerFunc(r.refreshTokenHandler.Handle)))
	mux.Handle("POST /api/authentication/v1/authentication/verify", loggingMiddleware(http.HandlerFunc(r.verifyTokenHandler.Handle)))
	mux.Handle("GET /api/authentication/v1/authentication/userinfo", loggingMiddleware(middleware.HandleWithValidateToken(r.userInfoHandler.Handle)))
}
