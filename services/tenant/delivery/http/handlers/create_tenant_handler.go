package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/tenant/delivery/http/dto"
	"services/tenant/service"
)

type CreateTenantHandler struct {
	service *service.TenantService
}

func NewCreateTenantHandler(service *service.TenantService) *CreateTenantHandler {
	return &CreateTenantHandler{service: service}
}

// Handle godoc
// @Summary Create tenant
// @Description Creates a new tenant.
// @Tags Tenant
// @Accept json
// @Produce json
// @Param request body dto.CreateTenantRequest true "Create tenant request"
// @Success 201 {object} dto.TenantResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/tenants [post]
func (h *CreateTenantHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateTenantRequest
	if err := json.Read(r, &req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
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

	json.Write(w, http.StatusCreated, response)
}
