package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/tenant/delivery/http/dto"
	"services/tenant/domain"
	"services/tenant/service"
)

type UpdateImageTenantHandler struct {
	service *service.TenantService
}

func NewUpdateImagesTenantHandler(service *service.TenantService) *UpdateImageTenantHandler {
	return &UpdateImageTenantHandler{service: service}
}

// Handle godoc
// @Summary Update tenant image
// @Description Updates tenant logo or banner image.
// @Tags Tenant
// @Accept multipart/form-data
// @Produce json
// @Param id path string true "Tenant ID"
// @Param type formData string true "Image type (logo or banner)"
// @Param image formData file true "Image file"
// @Success 200 {object} dto.TenantResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/tenants/{id}/image [put]
func (h *UpdateImageTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseMultipartForm(10 << 20); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	id := r.PathValue("id")
	imageType := domain.ImageTypeFromString(r.FormValue("type"))
	file, header, err := r.FormFile("image")
	if err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, "image file is required")
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

	json.Write(w, http.StatusOK, response)
}
