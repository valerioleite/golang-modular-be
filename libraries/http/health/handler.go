package health

import (
	"libraries/http/dto"
	"libraries/http/json"
	"net/http"
)

func Handler(moduleName string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		response := dto.HealthResponse{
			Status: "healthy",
			Module: moduleName,
		}

		json.Write(w, http.StatusOK, response)
	}
}
