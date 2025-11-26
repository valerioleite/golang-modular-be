package http

import (
	"errors"
	"libraries/http/dto"
	"libraries/http/json"
	"log/slog"
	"net/http"
	"time"
)

func HandleError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	errorMessage := err.Error()

	var httpErr dto.Error
	if errors.As(err, &httpErr) {
		status = httpErr.HTTPStatus()
		errorMessage = httpErr.Error()
	} else {
		slog.Error("Unexpected error occurred", "error", err)
	}

	HandleErrorWithStatus(w, status, errorMessage)
}

func HandleErrorWithStatus(w http.ResponseWriter, status int, error string) {
	response := dto.ErrorResponse{
		Timestamp: time.Now().UTC(),
		Errors:    []string{error},
	}

	json.Write(w, status, response)
}
