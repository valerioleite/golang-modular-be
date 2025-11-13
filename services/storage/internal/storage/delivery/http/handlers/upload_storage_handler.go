package handlers

import (
	"encoding/json"
	"mime/multipart"
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
	file, header, _ := r.FormFile("file")

	if file != nil {
		defer func(file multipart.File) {
			_ = file.Close()
		}(file)
	}

	storage, err := h.service.Upload(r.Context(), bucket, filename, file, header)
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
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		serverHttp.HandleError(w, err)
		return
	}
}
