package dto

import "time"

type CreateTenantRequest struct {
	Name string `json:"name"`
}

type UpdateTenantRequest struct {
	Name   string  `json:"name"`
	Logo   *string `json:"logo"`
	Banner *string `json:"banner"`
}

type TenantResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Logo   *string `json:"logo,omitempty"`
	Banner *string `json:"banner,omitempty"`
}

type ErrorResponse struct {
	Timestamp time.Time `json:"timestamp"`
	Errors    []string  `json:"errors"`
}
