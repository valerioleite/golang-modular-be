package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"mime/multipart"
	"net/http"
	"services/storage/delivery/http/dto"
	"services/storage/service"
)

type UploadStorageHandler struct {
	service *service.StorageService
}

func NewUploadStorageHandler(service *service.StorageService) *UploadStorageHandler {
	return &UploadStorageHandler{service: service}
}

// Handle godoc
// @Summary Upload file to storage
// @Description Uploads a file to the specified bucket.
// @Tags Storage
// @Accept multipart/form-data
// @Produce json
// @Param bucket formData string true "Bucket name"
// @Param filename formData string true "File name"
// @Param file formData file true "File to upload"
// @Success 201 {object} dto.StorageResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/files [post]
func (h *UploadStorageHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
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
		httpLib.HandleError(w, err)
		return
	}

	response := dto.StorageResponse{
		Bucket:   storage.Bucket,
		Filename: storage.Filename,
		Path:     storage.Path,
	}

	json.Write(w, http.StatusCreated, response)
}
