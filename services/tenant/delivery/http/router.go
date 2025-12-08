package http

import (
	"libraries/http/health"
	"libraries/http/middleware"
	"libraries/http/swagger"
	"net/http"
	_ "services/tenant/docs"
	"services/tenant/delivery/http/handlers"
	"services/tenant/service"
)

type Router struct {
	createHandler      *handlers.CreateTenantHandler
	getHandler         *handlers.GetTenantHandler
	listHandler        *handlers.ListTenantHandler
	updateHandler      *handlers.UpdateTenantHandler
	updateImageHandler *handlers.UpdateImageTenantHandler
	deleteHandler      *handlers.DeleteTenantHandler
}

func NewRouter(service *service.TenantService) *Router {
	return &Router{
		createHandler:      handlers.NewCreateTenantHandler(service),
		getHandler:         handlers.NewGetTenantHandler(service),
		listHandler:        handlers.NewListTenantHandler(service),
		updateHandler:      handlers.NewUpdateTenantHandler(service),
		updateImageHandler: handlers.NewUpdateImagesTenantHandler(service),
		deleteHandler:      handlers.NewDeleteTenantHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	loggingMiddleware := middleware.WithLogging("tenant")

	mux.HandleFunc("GET /api/tenant/swagger/", swagger.Handler("tenant"))

	r.setupActuatorResources(mux, loggingMiddleware)
	r.setupTenantResources(mux, loggingMiddleware)

	return mux
}

// healthCheck godoc
// @Summary Health check
// @Description Returns the health status of the tenant module.
// @Tags Actuator
// @Produce json
// @Success 200 {object} dto.HealthResponse
// @Router /v1/actuator/health [get]
func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	health.Handler("tenant")(w, req)
}

func (r *Router) setupActuatorResources(mux *http.ServeMux, loggingMiddleware func(http.Handler) http.Handler) {
	mux.Handle("GET /api/tenant/v1/actuator/health", loggingMiddleware(http.HandlerFunc(r.healthCheck)))
}

func (r *Router) setupTenantResources(mux *http.ServeMux, loggingMiddleware func(http.Handler) http.Handler) {
	mux.Handle("POST /api/tenant/v1/tenants", loggingMiddleware(http.HandlerFunc(r.createHandler.Handle)))
	mux.Handle("GET /api/tenant/v1/tenants", loggingMiddleware(http.HandlerFunc(r.listHandler.Handle)))
	mux.Handle("GET /api/tenant/v1/tenants/{id}", loggingMiddleware(http.HandlerFunc(r.getHandler.Handle)))
	mux.Handle("PUT /api/tenant/v1/tenants/{id}", loggingMiddleware(http.HandlerFunc(r.updateHandler.Handle)))
	mux.Handle("PUT /api/tenant/v1/tenants/{id}/image", loggingMiddleware(http.HandlerFunc(r.updateImageHandler.Handle)))
	mux.Handle("DELETE /api/tenant/v1/tenants/{id}", loggingMiddleware(http.HandlerFunc(r.deleteHandler.Handle)))
}
