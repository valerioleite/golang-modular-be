package swagger

import (
	"net/http"

	httpSwagger "github.com/swaggo/http-swagger"
)

func Handler(instanceName string) http.HandlerFunc {
	return func(w http.ResponseWriter, req *http.Request) {
		httpSwagger.Handler(httpSwagger.InstanceName(instanceName)).ServeHTTP(w, req)
	}
}
