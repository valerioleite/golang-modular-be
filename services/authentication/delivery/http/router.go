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
	userInfoHandler     *handlers.UserInfoHandler
}

func NewRouter(service *service.AuthenticationService) *Router {
	return &Router{
		loginHandler:        handlers.NewAuthorizeHandler(service),
		callbackGetHandler:  handlers.NewCallbackGetHandler(service),
		callbackPostHandler: handlers.NewCallbackPostHandler(service),
		refreshTokenHandler: handlers.NewRefreshTokenHandler(service),
		userInfoHandler:     handlers.NewUserInfoHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	requestContextMiddleware := middleware.WithRequestContext("authentication")

	mux.HandleFunc("GET /api/authentication/swagger/", swagger.Handler("authentication"))

	r.setupActuatorResources(mux, requestContextMiddleware)
	r.setupAuthenticationResources(mux, requestContextMiddleware)

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

func (r *Router) setupActuatorResources(mux *http.ServeMux, requestContextMiddleware func(http.Handler) http.Handler) {
	mux.Handle("GET /api/authentication/v1/actuator/health", requestContextMiddleware(http.HandlerFunc(r.healthCheck)))
}

func (r *Router) setupAuthenticationResources(mux *http.ServeMux, requestContextMiddleware func(http.Handler) http.Handler) {
	mux.Handle("GET /api/authentication/v1/authentication/authorize", requestContextMiddleware(http.HandlerFunc(r.loginHandler.Handle)))
	mux.Handle("GET /api/authentication/v1/authentication/callback", requestContextMiddleware(http.HandlerFunc(r.callbackGetHandler.Handle)))
	mux.Handle("POST /api/authentication/v1/authentication/callback", requestContextMiddleware(http.HandlerFunc(r.callbackPostHandler.Handle)))
	mux.Handle("POST /api/authentication/v1/authentication/refresh", requestContextMiddleware(http.HandlerFunc(r.refreshTokenHandler.Handle)))
	mux.Handle("GET /api/authentication/v1/authentication/userinfo", requestContextMiddleware(http.HandlerFunc(r.userInfoHandler.Handle)))
}
