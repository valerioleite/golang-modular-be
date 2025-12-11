package http

import (
	"libraries/http/health"
	"libraries/http/middleware"
	"libraries/http/swagger"
	"net/http"
	"services/user/delivery/http/handlers"
	_ "services/user/docs"
	"services/user/service"
)

type Router struct {
	createHandler *handlers.CreateUserHandler
	getHandler    *handlers.GetUserHandler
}

func NewRouter(service *service.UserService) *Router {
	return &Router{
		createHandler: handlers.NewCreateUserHandler(service),
		getHandler:    handlers.NewGetUserHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	requestContextMiddleware := middleware.WithRequestContext("user")

	mux.HandleFunc("GET /api/user/swagger/", swagger.Handler("user"))

	r.setupActuatorResources(mux, requestContextMiddleware)
	r.setupUserResources(mux, requestContextMiddleware)

	return mux
}

// healthCheck godoc
// @Summary Health check
// @Description Returns the health status of the user module.
// @Tags Actuator
// @Produce json
// @Success 200 {object} map[string]string
// @Router /v1/actuator/health [get]
func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	health.Handler("user")(w, req)
}

func (r *Router) setupActuatorResources(mux *http.ServeMux, requestContextMiddleware func(http.Handler) http.Handler) {
	mux.Handle("GET /api/user/v1/actuator/health", requestContextMiddleware(http.HandlerFunc(r.healthCheck)))
}

func (r *Router) setupUserResources(mux *http.ServeMux, requestContextMiddleware func(http.Handler) http.Handler) {
	mux.Handle("POST /api/user/v1/users", requestContextMiddleware(http.HandlerFunc(r.createHandler.Handle)))
	mux.Handle("GET /api/user/v1/users/sub/{sub}", requestContextMiddleware(http.HandlerFunc(r.getHandler.Handle)))
}
