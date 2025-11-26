package handlers

import (
	"log/slog"
	"net/http"
	"net/url"
	"services/authentication/internal/authentication/service"
	"strconv"
)

type CallbackGetHandler struct {
	service     *service.AuthenticationService
	frontendURL string
}

func NewCallbackGetHandler(service *service.AuthenticationService, frontendURL string) *CallbackGetHandler {
	return &CallbackGetHandler{
		service:     service,
		frontendURL: frontendURL,
	}
}

func (h *CallbackGetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	token, err := h.service.Callback(r.Context(), code, state)
	if err != nil {
		slog.Error("Failed to exchange code", "error", err)
		http.Redirect(w, r, h.frontendURL+"/callback?error="+url.QueryEscape(err.Error()), http.StatusFound)
		return
	}

	redirectURL := h.frontendURL + "/callback#access_token=" + url.QueryEscape(token.AccessToken) +
		"&id_token=" + url.QueryEscape(token.IDToken) +
		"&expires_in=" + strconv.FormatInt(token.ExpiresIn, 10) +
		"&token_type=" + url.QueryEscape(token.TokenType)

	if token.RefreshToken != "" {
		redirectURL += "&refresh_token=" + url.QueryEscape(token.RefreshToken)
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
}
