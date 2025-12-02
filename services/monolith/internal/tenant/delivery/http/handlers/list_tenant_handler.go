package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/monolith/internal/tenant/delivery/http/dto"
	"services/monolith/internal/tenant/service"
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
		httpLib.HandleError(w, err)
		return
	}

	response := make([]dto.TenantResponse, len(tenants))
	for i, tenant := range tenants {
		response[i] = dto.TenantResponse{
			ID:     tenant.ID.String(),
			Name:   tenant.Name,
			Logo:   tenant.Logo,
			Banner: tenant.Banner,
		}
	}

	json.Write(w, http.StatusOK, response)
}
