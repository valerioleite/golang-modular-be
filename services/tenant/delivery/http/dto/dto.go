package dto

type CreateTenantRequest struct {
	Name string `json:"name"`
} //@name CreateTenantRequest

type UpdateTenantRequest struct {
	Name   string  `json:"name"`
	Logo   *string `json:"logo"`
	Banner *string `json:"banner"`
} //@name UpdateTenantRequest

type TenantResponse struct {
	ID     string  `json:"id"`
	Name   string  `json:"name"`
	Logo   *string `json:"logo,omitempty"`
	Banner *string `json:"banner,omitempty"`
} //@name TenantResponse
