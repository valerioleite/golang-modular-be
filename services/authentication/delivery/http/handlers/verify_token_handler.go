package handlers

import (
	httpLib "libraries/http"
	"libraries/http/json"
	"net/http"
	"services/authentication/delivery/http/dto"
	"services/authentication/service"
)

type VerifyTokenHandler struct {
	service *service.AuthenticationService
}

func NewVerifyTokenHandler(service *service.AuthenticationService) *VerifyTokenHandler {
	return &VerifyTokenHandler{service: service}
}

// Handle godoc
// @Summary Verify ID token
// @Description Verifies ID token and returns user information.
// @Tags Authentication
// @Accept json
// @Produce json
// @Param request body dto.VerifyTokenRequest true "Token verification request"
// @Success 200 {object} dto.UserInfoResponse
// @Failure 400 {object} dto.ErrorResponse
// @Failure 401 {object} dto.ErrorResponse
// @Failure 500 {object} dto.ErrorResponse
// @Router /v1/authentication/verify [post]
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
		Subject:  userInfo.Subject,
		Email:    userInfo.Email,
		Name:     userInfo.Name,
		Username: &userInfo.PreferredUsername,
	}

	json.Write(w, http.StatusOK, response)
}
