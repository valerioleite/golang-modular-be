package dto

type UploadStorageRequest struct {
	Bucket   string `form:"bucket" binding:"required"`
	Filename string `form:"filename" binding:"required"`
}

type StorageResponse struct {
	ID       string `json:"id"`
	Bucket   string `json:"bucket"`
	Filename string `json:"filename"`
	Key      string `json:"key"`
	URL      string `json:"url,omitempty"`
}

