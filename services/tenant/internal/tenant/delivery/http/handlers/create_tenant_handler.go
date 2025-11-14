package handlers

import (
	"encoding/json"
	"net/http"
	httpLib "libraries/http"
	"services/tenant/internal/tenant/delivery/http/dto"
	"services/tenant/internal/tenant/service"
)

type CreateTenantHandler struct {
	service *service.TenantService
}

func NewCreateTenantHandler(service *service.TenantService) *CreateTenantHandler {
	return &CreateTenantHandler{service: service}
}

func (h *CreateTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTenantRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, []string{err.Error()})
		return
	}

	tenant, err := h.service.Create(r.Context(), req.Name, nil, nil)
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
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
