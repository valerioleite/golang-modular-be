package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/authentication/delivery/http/dto"
	"services/authentication/service"
)

type AuthorizeHandler struct {
	service *service.AuthenticationService
}

func NewAuthorizeHandler(service *service.AuthenticationService) *AuthorizeHandler {
	return &AuthorizeHandler{service: service}
}

// Handle godoc
// @Summary Start OAuth/OIDC authorization flow
// @Description Initiates the OAuth/OIDC authorization process and returns the authorization URL.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param redirect_uri query string true "Redirect URI after authentication"
// @Success 200 {object} dto.LoginResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/authentication/authorize [get]
func (h *AuthorizeHandler) Handle(w http.ResponseWriter, r *http.Request) {
	redirectURI := r.URL.Query().Get("redirect_uri")
	if redirectURI == "" {
		httpLib.HandleErrorWithStatus(w, http.StatusBadRequest, "redirect_uri is required")
		return
	}

	authURL, err := h.service.Authorize(r.Context(), redirectURI)
	if err != nil {
		httpLib.HandleError(w, err)
		return
	}

	response := dto.LoginResponse{
		AuthURL: authURL,
	}

	json.Write(w, http.StatusOK, response)
}
