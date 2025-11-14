package handlers

import (
	"encoding/json"
	httpLib "libraries/http"
	"net/http"
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

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
