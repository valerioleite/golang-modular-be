package http

import (
	"net/http"
	"services/tenant/internal/tenant/delivery/http/handlers"
	"services/tenant/internal/tenant/service"
)

type Router struct {
	createHandler       *handlers.CreateTenantHandler
	getHandler          *handlers.GetTenantHandler
	listHandler         *handlers.ListTenantHandler
	updateHandler       *handlers.UpdateTenantHandler
	updateImagesHandler *handlers.UpdateImagesTenantHandler
	deleteHandler       *handlers.DeleteTenantHandler
}

func NewRouter(service *service.TenantService) *Router {
	return &Router{
		createHandler:       handlers.NewCreateTenantHandler(service),
		getHandler:          handlers.NewGetTenantHandler(service),
		listHandler:         handlers.NewListTenantHandler(service),
		updateHandler:       handlers.NewUpdateTenantHandler(service),
		updateImagesHandler: handlers.NewUpdateImagesTenantHandler(service),
		deleteHandler:       handlers.NewDeleteTenantHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /tenants", r.createHandler.Handle)
	mux.HandleFunc("GET /tenants", r.listHandler.Handle)
	mux.HandleFunc("GET /tenants/{id}", r.getHandler.Handle)
	mux.HandleFunc("PUT /tenants/{id}", r.updateHandler.Handle)
	mux.HandleFunc("PUT /tenants/{id}/image", r.updateImagesHandler.Handle)
	mux.HandleFunc("DELETE /tenants/{id}", r.deleteHandler.Handle)

	return mux
}
