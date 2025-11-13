package http

import (
	"net/http"
	"services/storage/internal/storage/delivery/http/handlers"
	"services/storage/internal/storage/service"
)

type Router struct {
	uploadHandler *handlers.UploadStorageHandler
}

func NewRouter(service *service.StorageService) *Router {
	return &Router{
		uploadHandler: handlers.NewUploadStorageHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("POST /storage", r.Upload)

	return mux
}

func (r *Router) Upload(w http.ResponseWriter, req *http.Request) {
	r.uploadHandler.Handle(w, req)
}

