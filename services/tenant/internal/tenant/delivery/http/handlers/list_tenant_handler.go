package handlers

import (
	"encoding/json"
	"net/http"
	serverHttp "services/tenant/internal/server/http"
	"services/tenant/internal/tenant/delivery/http/dto"
	"services/tenant/internal/tenant/service"
)

type ListTenantHandler struct {
	service *service.TenantService
}

func NewListTenantHandler(service *service.TenantService) *ListTenantHandler {
	return &ListTenantHandler{service: service}
}

func (h *ListTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	tenants, err := h.service.List(r.Context())
	if err != nil {
		serverHttp.HandleError(w, err)
		return
	}

	responses := make([]dto.TenantResponse, len(tenants))
	for i, tenant := range tenants {
		responses[i] = dto.TenantResponse{
			ID:     tenant.ID.String(),
			Name:   tenant.Name,
			Logo:   tenant.Logo,
			Banner: tenant.Banner,
		}
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(responses)
}
