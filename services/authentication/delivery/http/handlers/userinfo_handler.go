package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
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
	authHeader := r.Header.Get("Authorization")
	if authHeader == "" {
		httpLib.HandleErrorWithStatus(w, http.StatusUnauthorized, "missing authorization header")
		return
	}

	accessToken := authHeader
	if len(authHeader) > 7 && authHeader[:7] == "Bearer " {
		accessToken = authHeader[7:]
	}

	user, err := h.service.GetUserInfo(r.Context(), accessToken)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.UserInfoResponse{
		Subject:  user.Sub,
		Email:    user.Email,
		Name:     user.Name,
		Username: user.Username,
	}

	json.Write(w, http.StatusOK, response)
}
