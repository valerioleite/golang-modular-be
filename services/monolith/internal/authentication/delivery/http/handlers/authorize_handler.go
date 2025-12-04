package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/monolith/internal/authentication/delivery/http/dto"
	"services/monolith/internal/authentication/service"
)

type AuthorizeHandler struct {
	service *service.AuthenticationService
}

func NewAuthorizeHandler(service *service.AuthenticationService) *AuthorizeHandler {
	return &AuthorizeHandler{service: service}
}

func (h *AuthorizeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect_uri")
	if redirectURI == "" {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, "redirect_uri is required")
		return
	}

	authURL, err := h.service.Login(r.Context(), redirectURI)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.LoginResponse{
		AuthURL: authURL,
	}

	json.Write(w, http.StatusOK, response)
}
