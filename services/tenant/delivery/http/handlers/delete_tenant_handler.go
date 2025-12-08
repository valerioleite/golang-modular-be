package handlers

import (
	httpLib "libraries/http"
	"net/http"
	"services/tenant/service"
)

type DeleteTenantHandler struct {
	service *service.TenantService
}

func NewDeleteTenantHandler(service *service.TenantService) *DeleteTenantHandler {
	return &DeleteTenantHandler{service: service}
}

// Handle godoc
// @Summary Delete tenant
// @Description Deletes a tenant by ID.
// @Tags Tenant
// @Param id path string true "Tenant ID"
// @Success 204 "No Content"
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/tenants/{id} [delete]
func (h *DeleteTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	if err := h.service.Delete(r.Context(), id); err != nil {
		httpLib.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
