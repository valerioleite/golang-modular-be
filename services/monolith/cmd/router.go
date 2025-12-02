package main

import (
	"libraries/domain"
	httpLib "libraries/http"
	"net/http"
	authHttp "services/monolith/internal/authentication/delivery/http"
	storageHttp "services/monolith/internal/storage/delivery/http"
	tenantHttp "services/monolith/internal/tenant/delivery/http"
)

type Router struct {
	authRouter    *authHttp.Router
	tenantRouter  *tenantHttp.Router
	storageRouter *storageHttp.Router
}

func NewRouter(authRouter *authHttp.Router, tenantRouter *tenantHttp.Router, storageRouter *storageHttp.Router) *Router {
	return &Router{
		authRouter:    authRouter,
		tenantRouter:  tenantRouter,
		storageRouter: storageRouter,
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	authMux := r.authRouter.SetupRoutes()
	tenantMux := r.tenantRouter.SetupRoutes()
	storageMux := r.storageRouter.SetupRoutes()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		muxes := []*http.ServeMux{authMux, tenantMux, storageMux}
		for _, m := range muxes {
			if handlerFound(m, req) {
				m.ServeHTTP(w, req)
				return
			}
		}

		httpLib.HandleError(w, domain.NewNotFoundError("route not found"))
	})

	return mux
}
