package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"libraries/http/middleware"
	"net/http"
	"services/authentication/delivery/http/dto"
	"services/authentication/service"
)

type UserInfoHandler struct {
	service *service.AuthenticationService
}

func NewUserInfoHandler(service *service.AuthenticationService) *UserInfoHandler {
	return &UserInfoHandler{service: service}
}

// Handle godoc
// @Summary Get user information
// @Description Retrieves user information from access token.
// @Tags Authentication
// @Accept json
// @Produce json
// @Success 200 {object} dto.UserInfoResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/authentication/userinfo [get]
// @Security BearerAuth
func (h *UserInfoHandler) Handle(w http.ResponseWriter, r *http.Request) {
	accessToken := middleware.GetAccessToken(r.Context())

	user, err := h.service.GetUserInfo(r.Context(), accessToken)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.UserInfoResponse{
		Sub:       user.Sub,
		Email:     user.Email,
		Username:  user.Username,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}

	json.Write(w, http.StatusOK, response)
}
