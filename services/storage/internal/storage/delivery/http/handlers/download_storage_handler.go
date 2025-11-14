package handlers

import (
	"io"
	"net/http"
	serverHttp "services/storage/internal/server/http"
	"services/storage/internal/storage/service"
)

type DownloadStorageHandler struct {
	service *service.StorageService
}

func NewDownloadStorageHandler(service *service.StorageService) *DownloadStorageHandler {
	return &DownloadStorageHandler{service: service}
}

func (h *DownloadStorageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	bucket := r.PathValue("bucket")
	filename := r.PathValue("filename")

	if bucket == "" {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"bucket is required"})
		return
	}

	if filename == "" {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"filename is required"})
		return
	}

	reader, err := h.service.Download(r.Context(), bucket, filename)
	if err != nil {
		serverHttp.HandleError(w, err)
		return
	}
	defer reader.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, reader)
	if err != nil {
		serverHttp.HandleError(w, err)
		return
	}
}
