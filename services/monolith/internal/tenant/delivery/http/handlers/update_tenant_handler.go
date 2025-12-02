package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/monolith/internal/tenant/delivery/http/dto"
	"services/monolith/internal/tenant/service"
)

type UpdateTenantHandler struct {
	service *service.TenantService
}

func NewUpdateTenantHandler(service *service.TenantService) *UpdateTenantHandler {
	return &UpdateTenantHandler{service: service}
}

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
