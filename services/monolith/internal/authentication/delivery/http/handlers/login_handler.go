package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/monolith/internal/authentication/delivery/http/dto"
	"services/monolith/internal/authentication/service"
)

type LoginHandler struct {
	service *service.AuthenticationService
}

func NewLoginHandler(service *service.AuthenticationService) *LoginHandler {
	return &LoginHandler{service: service}
}

func (h *LoginHandler) Handle(w http.ResponseWriter, r *http.Request) {
	var req dto.LoginRequest
	if err := json.Read(r, &req); err != nil {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, err.Error())
		return
	}

	authURL, state, err := h.service.Login(r.Context(), req.RedirectURI)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.LoginResponse{
		AuthURL: authURL,
		State:   state,
	}

	json.Write(w, http.StatusOK, response)
}
