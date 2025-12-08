package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/authentication/delivery/http/dto"
	"services/authentication/service"
)

type CallbackPostHandler struct {
	service *service.AuthenticationService
}

func NewCallbackPostHandler(service *service.AuthenticationService) *CallbackPostHandler {
	return &CallbackPostHandler{service: service}
}

// Handle godoc
// @Summary OAuth callback handler (API)
// @Description Exchanges authorization code for tokens via API request.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.CallbackRequest true "Callback request with code and state"
// @Success 200 {object} dto.TokenResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/authentication/callback [post]
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
