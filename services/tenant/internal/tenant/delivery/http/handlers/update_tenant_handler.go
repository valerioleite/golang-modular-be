package handlers

import (
	"encoding/json"
	"net/http"
	serverHttp "services/tenant/internal/server/http"
	"services/tenant/internal/tenant/delivery/http/dto"
	"services/tenant/internal/tenant/service"
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
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	tenant, err := h.service.Update(r.Context(), id, req.Name, req.Logo, req.Banner)
	if err != nil {
		serverHttp.HandleError(w, err)
		return
	}

	response := dto.TenantResponse{
		ID:     tenant.ID.String(),
		Name:   tenant.Name,
		Logo:   tenant.Logo,
		Banner: tenant.Banner,
	}

	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(response) //TODO add in http library
	if err != nil {
		serverHttp.HandleError(w, err)
	}
}
