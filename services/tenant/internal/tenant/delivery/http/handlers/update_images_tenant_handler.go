package handlers

import (
	"encoding/json"
	httpLib "libraries/http"
	"net/http"
	"services/tenant/internal/tenant/delivery/http/dto"
	"services/tenant/internal/tenant/domain"
	"services/tenant/internal/tenant/service"
)

type UpdateImagesTenantHandler struct {
	service *service.TenantService
}

func NewUpdateImagesTenantHandler(service *service.TenantService) *UpdateImagesTenantHandler {
	return &UpdateImagesTenantHandler{service: service}
}

func (h *UpdateImagesTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"failed to parse multipart form"})
		return
	}

	id := r.PathValue("id")
	imageType := domain.ImageTypeFromString(r.FormValue("type"))
	file, header, err := r.FormFile("image")
	if err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, []string{"image file is required"})
		return
	}
	
	defer func() {
		if file != nil {
			_ = file.Close()
		}
	}()

	tenant, err := h.service.UpdateImage(r.Context(), id, imageType, file, header)
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
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(response)
	if err != nil {
		httpLib.HandleError(w, err)
	}
}
