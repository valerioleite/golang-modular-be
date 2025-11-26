package http

import (
	"net/http"

	"github.com/rs/cors"
)

func NewCORSHandler(handler http.Handler) http.Handler {
	c := cors.New(cors.Options{
		AllowedOrigins: []string{"*"},
		AllowedMethods: []string{
			http.MethodGet,
			http.MethodPost,
			http.MethodPut,
			http.MethodDelete,
			http.MethodOptions,
			http.MethodPatch,
		},
		AllowedHeaders: []string{
			"Content-Type",
			"Authorization",
			"Accept",
			"Origin",
			"X-Requested-With",
		},
		ExposedHeaders:   []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           3600,
	})

	return c.Handler(handler)
}

func NewCORSHandlerWithOptions(options cors.Options) func(http.Handler) http.Handler {
	c := cors.New(options)
	return c.Handler
}

