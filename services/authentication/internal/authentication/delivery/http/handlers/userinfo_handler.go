package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/authentication/internal/authentication/delivery/http/dto"
	"services/authentication/internal/authentication/service"
)

type UserInfoHandler struct {
	service *service.AuthenticationService
}

func NewUserInfoHandler(service *service.AuthenticationService) *UserInfoHandler {
	return &UserInfoHandler{service: service}
}

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

	userInfo, err := h.service.GetUserInfo(r.Context(), accessToken)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.UserInfoResponse{
		Subject:           userInfo.Subject,
		Email:             userInfo.Email,
		EmailVerified:     userInfo.EmailVerified,
		Name:              userInfo.Name,
		PreferredUsername: userInfo.PreferredUsername,
		Picture:           userInfo.Picture,
	}

	json.Write(w, http.StatusOK, response)
}
