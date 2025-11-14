package dto

type UploadStorageRequest struct {
	Bucket   string `form:"bucket" binding:"required"`
	Filename string `form:"filename" binding:"required"`
}

type StorageResponse struct {
	Bucket   string `json:"bucket"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
}
