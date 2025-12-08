package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/tenant/delivery/http/dto"
	"services/tenant/service"
)

type ListTenantHandler struct {
	service *service.TenantService
}

func NewListTenantHandler(service *service.TenantService) *ListTenantHandler {
	return &ListTenantHandler{service: service}
}

// Handle godoc
// @Summary List all tenants
// @Description Retrieves all tenants.
// @Tags Tenant
// @Produce json
// @Success 200 {array} dto.TenantResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/tenants [get]
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
