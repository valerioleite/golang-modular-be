package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/user/delivery/http/dto"
	"services/user/service"
)

type CreateUserHandler struct {
	service *service.UserService
}

func NewCreateUserHandler(service *service.UserService) *CreateUserHandler {
	return &CreateUserHandler{service: service}
}

// Handle godoc
// @Summary Create user
// @Description Creates a new user.
// @Tags User
// @Accept json
// @Produce json
// @Param request body dto.CreateUserRequest true "Create user request"
// @Success 201 {object} dto.UserResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 409 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/users [post]
func (h *CreateUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.CreateUserRequest
	if err := json.Read(r, &req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	// TODO: Get createdBy from authentication context
	createdBy := "system"

	user, err := h.service.Create(r.Context(), createdBy, req.Sub, req.Email, req.Name, req.Username)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.UserResponse{
		ID:        user.ID.String(),
		CreatedBy: user.CreatedBy,
		CreatedAt: user.CreatedAt.Format("2006-01-02T15:04:05Z07:00"),
		UpdatedBy: user.UpdatedBy,
		UpdatedAt: user.UpdatedAt.Format("2006-01-02T15:04:05Z07:00"),
		Sub:       user.Sub,
		Email:     user.Email,
		Name:      user.Name,
		Username:  user.Username,
	}

	json.Write(w, http.StatusCreated, response)
}
