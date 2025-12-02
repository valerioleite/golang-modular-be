package handlers

import (
	"log/slog"
	"net/http"
	"net/url"
	"services/monolith/internal/authentication/service"
	"strconv"
)

type CallbackGetHandler struct {
	service *service.AuthenticationService
}

func NewCallbackGetHandler(service *service.AuthenticationService) *CallbackGetHandler {
	return &CallbackGetHandler{
		service: service,
	}
}

func (h *CallbackGetHandler) Handle(w http.ResponseWriter, r *http.Request) {
	code := r.URL.Query().Get("code")
	state := r.URL.Query().Get("state")

	redirectURI, err := h.service.GetRedirectURI(state)
	if err != nil {
		slog.Error("Invalid state", "error", err)
		http.Redirect(w, r, redirectURI+"?error=invalid_state", http.StatusFound)
		return
	}

	token, err := h.service.Callback(r.Context(), code, state)
	if err != nil {
		slog.Error("Failed to exchange code", "error", err)
		http.Redirect(w, r, redirectURI+"?error="+url.QueryEscape(err.Error()), http.StatusFound)
		return
	}

	redirectURL := redirectURI +
		"#access_token=" + url.QueryEscape(token.AccessToken) +
		"&id_token=" + url.QueryEscape(token.IDToken) +
		"&expires_in=" + strconv.FormatInt(token.ExpiresIn, 10) +
		"&token_type=" + url.QueryEscape(token.TokenType)

	if token.RefreshToken != "" {
		redirectURL += "&refresh_token=" + url.QueryEscape(token.RefreshToken)
	}

	http.Redirect(w, r, redirectURL, http.StatusFound)
}
