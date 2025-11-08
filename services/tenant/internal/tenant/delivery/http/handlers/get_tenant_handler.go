package handlers

import (
	"encoding/json"
	"net/http"
	serverHttp "services/tenant/internal/server/http"
	"services/tenant/internal/tenant/delivery/http/dto"
	"services/tenant/internal/tenant/service"
)

type GetTenantHandler struct {
	service *service.TenantService
}

func NewGetTenantHandler(service *service.TenantService) *GetTenantHandler {
	return &GetTenantHandler{service: service}
}

func (h *GetTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		serverHttp.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"id is required"})
		return
	}

	tenant, err := h.service.Get(r.Context(), id)
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
	json.NewEncoder(w).Encode(response)
}
