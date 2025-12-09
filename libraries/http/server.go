package http

import (
	"libraries/http/cors"
	"log/slog"
	"net/http"
	"os"
)

type Router interface {
	SetupRoutes() *http.ServeMux
}

type Server struct {
	mux     *http.ServeMux
	handler http.Handler
	port    string
}

func NewServer(router Router) *Server {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	mux := router.SetupRoutes()
	handler := cors.NewCORSHandler(mux)

	return &Server{
		mux:     mux,
		port:    port,
		handler: handler,
	}
}

func (s *Server) Start() error {
	slog.Info("Server starting", "port", s.port)
	if err := http.ListenAndServe(":"+s.port, s.handler); err != nil {
		return err
	}
	
	return nil
}
