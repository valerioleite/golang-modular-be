package http

import (
	"libraries/http/health"
	"libraries/http/middleware"
	"libraries/http/swagger"
	"net/http"
	_ "services/storage/docs"
	"services/storage/delivery/http/handlers"
	"services/storage/service"
)

type Router struct {
	uploadHandler   *handlers.UploadStorageHandler
	downloadHandler *handlers.DownloadStorageHandler
}

func NewRouter(service *service.StorageService) *Router {
	return &Router{
		uploadHandler:   handlers.NewUploadStorageHandler(service),
		downloadHandler: handlers.NewDownloadStorageHandler(service),
	}
}

func (r *Router) SetupRoutes() *http.ServeMux {
	mux := http.NewServeMux()

	loggingMiddleware := middleware.WithLogging("storage")

	mux.HandleFunc("GET /api/storage/swagger/", swagger.Handler("storage"))

	r.setupActuatorResources(mux, loggingMiddleware)
	r.setupStorageResources(mux, loggingMiddleware)

	return mux
}

// healthCheck godoc
// @Summary Health check
// @Description Returns the health status of the storage module.
// @Tags Actuator
// @Produce json
// @Success 200 {object} dto.HealthResponse
// @Router /v1/actuator/health [get]
func (r *Router) healthCheck(w http.ResponseWriter, req *http.Request) {
	health.Handler("storage")(w, req)
}

func (r *Router) setupActuatorResources(mux *http.ServeMux, loggingMiddleware func(http.Handler) http.Handler) {
	mux.Handle("GET /api/storage/v1/actuator/health", loggingMiddleware(http.HandlerFunc(r.healthCheck)))
}

func (r *Router) setupStorageResources(mux *http.ServeMux, loggingMiddleware func(http.Handler) http.Handler) {
	mux.Handle("POST /api/storage/v1/files", loggingMiddleware(http.HandlerFunc(r.uploadHandler.Handle)))
	mux.Handle("GET /api/storage/v1/files/{bucket}/{filename}", loggingMiddleware(http.HandlerFunc(r.downloadHandler.Handle)))
}

