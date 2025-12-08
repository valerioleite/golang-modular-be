package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/tenant/delivery/http/dto"
	"services/tenant/service"
)

type GetTenantHandler struct {
	service *service.TenantService
}

func NewGetTenantHandler(service *service.TenantService) *GetTenantHandler {
	return &GetTenantHandler{service: service}
}

// Handle godoc
// @Summary Get tenant by ID
// @Description Retrieves a tenant by its ID.
// @Tags Tenant
// @Produce json
// @Param id path string true "Tenant ID"
// @Success 200 {object} dto.TenantResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/tenants/{id} [get]
func (h *GetTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	tenant, err := h.service.Get(r.Context(), id)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.TenantResponse{
		ID:     tenant.ID.String(),
		Name:   tenant.Name,
		Logo:   tenant.Logo,
		Banner: tenant.Banner,
	}

	json.Write(w, http.StatusOK, response)
}
