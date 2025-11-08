package http

import (
	"net/http"
	"services/tenant/internal/tenant/delivery/http/handlers"
	"services/tenant/internal/tenant/service"
)

type Router struct {
	createHandler *handlers.CreateTenantHandler
	getHandler    *handlers.GetTenantHandler
	listHandler   *handlers.ListTenantHandler
	updateHandler *handlers.UpdateTenantHandler
	deleteHandler *handlers.DeleteTenantHandler
}

func NewRouter(service *service.TenantService) *Router {
	return &Router{
		createHandler: handlers.NewCreateTenantHandler(service),
		getHandler:    handlers.NewGetTenantHandler(service),
		listHandler:   handlers.NewListTenantHandler(service),
		updateHandler: handlers.NewUpdateTenantHandler(service),
		deleteHandler: handlers.NewDeleteTenantHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /tenants", r.Create)
	mux.HandleFunc("GET /tenants", r.List)
	mux.HandleFunc("GET /tenants/{id}", r.Get)
	mux.HandleFunc("PUT /tenants/{id}", r.Update)
	mux.HandleFunc("DELETE /tenants/{id}", r.Delete)

	return mux
}

func (r *Router) Create(w http.ResponseWriter, req *http.Request) {
	r.createHandler.Handle(w, req)
}

func (r *Router) Get(w http.ResponseWriter, req *http.Request) {
	r.getHandler.Handle(w, req)
}

func (r *Router) List(w http.ResponseWriter, req *http.Request) {
	r.listHandler.Handle(w, req)
}

func (r *Router) Update(w http.ResponseWriter, req *http.Request) {
	r.updateHandler.Handle(w, req)
}

func (r *Router) Delete(w http.ResponseWriter, req *http.Request) {
	r.deleteHandler.Handle(w, req)
}
