package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/authentication/internal/authentication/delivery/http/dto"
	"services/authentication/internal/authentication/service"
)

type RefreshTokenHandler struct {
	service *service.AuthenticationService
}

func NewRefreshTokenHandler(service *service.AuthenticationService) *RefreshTokenHandler {
	return &RefreshTokenHandler{service: service}
}

func (h *RefreshTokenHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.RefreshTokenRequest
	if err := json.Read(r, &req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.RefreshToken(r.Context(), req.RefreshToken)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.TokenResponse{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		IDToken:      token.IDToken,
		ExpiresIn:    token.ExpiresIn,
		TokenType:    token.TokenType,
	}

	json.Write(w, http.StatusOK, response)
}

