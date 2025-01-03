package server

import (
	"net/http"

	"github.com/praveenmahasena/libsys/internal/server/handlers"
)

type Server struct {
	Addr string
}

func New(p string) *Server {
	return &Server{
		p,
	}
}

func (s *Server) Run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("GET /",handlers.Home)
	mux.HandleFunc("GET /add",handlers.Add)
	return http.ListenAndServe(s.Addr, mux)
}
