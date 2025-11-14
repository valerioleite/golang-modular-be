package http

import (
	"encoding/json"
	"errors"
	"net/http"
	"time"
)

func HandleError(w http.ResponseWriter, err error) {
	statusCode := http.StatusInternalServerError
	errorMessage := err.Error()

	var httpErr Error
	if errors.As(err, &httpErr) {
		statusCode = httpErr.HTTPStatus()
		errorMessage = httpErr.Error()
	}

	HandleErrorWithStatus(w, statusCode, []string{errorMessage})
}

func HandleErrorWithStatus(w http.ResponseWriter, statusCode int, errors []string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	response := ErrorResponse{
		Timestamp: time.Now().UTC(),
		Errors:    errors,
	}

	_ = json.NewEncoder(w).Encode(response)
}
