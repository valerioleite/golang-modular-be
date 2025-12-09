package main

import (
	"libraries/domain"
	httpLib "libraries/http"
	"net/http"
	authHttp "services/authentication/delivery/http"
	storageHttp "services/storage/delivery/http"
	tenantHttp "services/tenant/delivery/http"
	userHttp "services/user/delivery/http"
)

type Router struct {
	authRouter    *authHttp.Router
	tenantRouter  *tenantHttp.Router
	storageRouter *storageHttp.Router
	userRouter    *userHttp.Router
}

func NewRouter(authRouter *authHttp.Router, tenantRouter *tenantHttp.Router, storageRouter *storageHttp.Router, userRouter *userHttp.Router) *Router {
	return &Router{
		authRouter:    authRouter,
		tenantRouter:  tenantRouter,
		storageRouter: storageRouter,
		userRouter:    userRouter,
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	authMux := r.authRouter.SetupRoutes()
	tenantMux := r.tenantRouter.SetupRoutes()
	storageMux := r.storageRouter.SetupRoutes()
	userMux := r.userRouter.SetupRoutes()

	mux.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		muxes := []*http.ServeMux{authMux, tenantMux, storageMux, userMux}
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
