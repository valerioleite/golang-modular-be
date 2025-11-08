package http

import (
	"encoding/json"
	"net/http"
	"time"
)

func HandleError(w http.ResponseWriter, err error) {
	var statusCode int
	var errorMessage string

	if httpErr, ok := err.(HTTPError); ok {
		statusCode = httpErr.HTTPStatus()
		errorMessage = httpErr.Error()
	} else {
		statusCode = http.StatusInternalServerError
		errorMessage = err.Error()
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
	json.NewEncoder(w).Encode(response)
}
