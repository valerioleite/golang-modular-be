package handlers

import (
	"net/http"
	httpLib "libraries/http"
	"services/tenant/internal/tenant/service"
)

type DeleteTenantHandler struct {
	service *service.TenantService
}

func NewDeleteTenantHandler(service *service.TenantService) *DeleteTenantHandler {
	return &DeleteTenantHandler{service: service}
}

func (h *DeleteTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")
	if id == "" {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"id is required"})
		return
	}

	if err := h.service.Delete(r.Context(), id); err != nil {
		httpLib.HandleError(w, err)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
