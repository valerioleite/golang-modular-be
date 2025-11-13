package handlers

import (
	"encoding/json"
	"net/http"
	serverHttp "services/storage/internal/server/http"
	"services/storage/internal/storage/delivery/http/dto"
	"services/storage/internal/storage/service"
)

type UploadStorageHandler struct {
	service *service.StorageService
}

func NewUploadStorageHandler(service *service.StorageService) *UploadStorageHandler {
	return &UploadStorageHandler{service: service}
}

func (h *UploadStorageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"failed to parse multipart form"})
		return
	}

	bucket := r.FormValue("bucket")
	filename := r.FormValue("filename")

	if bucket == "" {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"bucket is required"})
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"file is required"})
		return
	}
	defer file.Close()

	if filename == "" {
		filename = header.Filename
	}

	if filename == "" {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"filename is required"})
		return
	}

	storage, err := h.service.Upload(r.Context(), bucket, filename, file)
	if err != nil {
		serverHttp.HandleError(w, err)
		return
	}

	response := dto.StorageResponse{
		ID:       storage.ID.String(),
		Bucket:   storage.Bucket,
		Filename: storage.Filename,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
