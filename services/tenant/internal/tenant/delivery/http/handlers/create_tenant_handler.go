package handlers

import (
	"encoding/json"
	httpLib "libraries/http"
	"net/http"
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

	tenant, err := h.service.Create(r.Context(), req.Name)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.TenantResponse{
		ID:   tenant.ID.String(),
		Name: tenant.Name,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpLib.HandleError(w, err)
	}
}
