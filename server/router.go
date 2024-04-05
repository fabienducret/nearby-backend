package server

import (
	"github.com/go-chi/chi"
)

func (s *Server) Router(c Controllers) *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", c.Health)

	return r
}
