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

	mux.HandleFunc("/",handlers.Home)

	return http.ListenAndServe(s.Addr, mux)
}
