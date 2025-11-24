package http

import (
	"errors"
	json "libraries/http/json"
	"log/slog"
	"net/http"
	"time"
)

func HandleError(w http.ResponseWriter, err error) {
	status := http.StatusInternalServerError
	errorMessage := err.Error()

	var httpErr Error
	if errors.As(err, &httpErr) {
		status = httpErr.HTTPStatus()
		errorMessage = httpErr.Error()
	} else {
		slog.Error("Unexpected error occurred", "error", err)
	}

	HandleErrorWithStatus(w, status, errorMessage)
}

func HandleErrorWithStatus(w http.ResponseWriter, status int, errors string) {
	response := ErrorResponse{
		Timestamp: time.Now().UTC(),
		Errors:    []string{errors},
	}

	json.Write(w, status, response)
}
