package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/monolith/internal/tenant/delivery/http/dto"
	"services/monolith/internal/tenant/service"
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

	json.Write(w, http.StatusOK, response)
}
