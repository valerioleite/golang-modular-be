package server

import (
	"log/slog"
	"net/http"
	"os"
	tenantHttp "services/tenant/internal/tenant/delivery/http"
)

type Server struct {
	mux    *http.ServeMux
	router *tenantHttp.Router
	port   string
}

func NewServer(router *tenantHttp.Router) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := router.SetupRoutes()

	return &Server{
		mux:    mux,
		router: router,
		port:   port,
	}
}

func (s *Server) Start() error {
	slog.Info("Server starting.", "port", s.port)
	if err := http.ListenAndServe(":"+s.port, s.mux); err != nil {
		return err
	}
	return nil
}
