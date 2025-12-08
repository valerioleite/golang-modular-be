package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/tenant/delivery/http/dto"
	"services/tenant/service"
)

type UpdateTenantHandler struct {
	service *service.TenantService
}

func NewUpdateTenantHandler(service *service.TenantService) *UpdateTenantHandler {
	return &UpdateTenantHandler{service: service}
}

// Handle godoc
// @Summary Update tenant
// @Description Updates an existing tenant.
// @Tags Tenant
// @Accept json
// @Produce json
// @Param id path string true "Tenant ID"
// @Param request body dto.UpdateTenantRequest true "Update tenant request"
// @Success 200 {object} dto.TenantResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/tenants/{id} [put]
func (h *UpdateTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	var req dto.UpdateTenantRequest
	if err := json.Read(r, &req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	tenant, err := h.service.Update(r.Context(), id, req.Name)
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
