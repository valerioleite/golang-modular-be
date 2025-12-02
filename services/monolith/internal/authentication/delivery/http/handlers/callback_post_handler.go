package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/monolith/internal/authentication/delivery/http/dto"
	"services/monolith/internal/authentication/service"
)

type CallbackPostHandler struct {
	service *service.AuthenticationService
}

func NewCallbackPostHandler(service *service.AuthenticationService) *CallbackPostHandler {
	return &CallbackPostHandler{service: service}
}

func (h *CallbackPostHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.CallbackRequest
	if err := json.Read(r, &req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	token, err := h.service.Callback(r.Context(), req.Code, req.State)
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

