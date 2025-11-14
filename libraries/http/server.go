package http

import (
	"log/slog"
	"net/http"
	"os"
)

type Router interface {
	SetupRoutes() *http.ServeMux
}

type Server struct {
	mux  *http.ServeMux
	port string
}

func NewServer(router Router) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := router.SetupRoutes()

	return &Server{
		mux:  mux,
		port: port,
	}
}

func (s *Server) Start() error {
	slog.Info("Server starting.", "port", s.port)
	if err := http.ListenAndServe(":"+s.port, s.mux); err != nil {
		return err
	}
	return nil
}

