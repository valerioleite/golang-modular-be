package handlers

import (
	"io"
	httpLib "libraries/http"
	"net/http"
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
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"bucket is required"})
		return
	}

	if filename == "" {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"filename is required"})
		return
	}

	reader, err := h.service.Download(r.Context(), bucket, filename)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}
	defer reader.Close()

	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Disposition", "attachment; filename=\""+filename+"\"")
	w.WriteHeader(http.StatusOK)

	_, err = io.Copy(w, reader)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}
}
