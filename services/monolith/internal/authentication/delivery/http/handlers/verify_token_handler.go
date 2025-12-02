package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/monolith/internal/authentication/delivery/http/dto"
	"services/monolith/internal/authentication/service"
)

type VerifyTokenHandler struct {
	service *service.AuthenticationService
}

func NewVerifyTokenHandler(service *service.AuthenticationService) *VerifyTokenHandler {
	return &VerifyTokenHandler{service: service}
}

func (h *VerifyTokenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.VerifyTokenRequest
	if err := json.Read(r, &req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	userInfo, err := h.service.VerifyToken(r.Context(), req.Token)
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

