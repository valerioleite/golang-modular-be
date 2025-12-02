package dto

type UploadResponse struct {
	Bucket   string `json:"bucket"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
}
