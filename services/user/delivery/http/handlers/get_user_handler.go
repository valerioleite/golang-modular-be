package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/user/delivery/http/dto"
	"services/user/service"
)

type GetUserHandler struct {
	service *service.UserService
}

func NewGetUserHandler(service *service.UserService) *GetUserHandler {
	return &GetUserHandler{service: service}
}

// Handle godoc
// @Summary Get user by sub
// @Description Retrieves a user by their sub (subject identifier).
// @Tags User
// @Produce json
// @Param sub path string true "User sub"
// @Success 200 {object} dto.UserResponse
// @Failure 404 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/users/sub/{sub} [get]
func (h *GetUserHandler) Handle(w http.ResponseWriter, r *http.Request) {
	sub := r.PathValue("sub")

	user, err := h.service.GetBySub(r.Context(), sub)
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

	json.Write(w, http.StatusOK, response)
}
