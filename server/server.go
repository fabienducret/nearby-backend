package server

import (
	"fmt"
	"net/http"
)

type Controllers struct {
	Health       http.HandlerFunc
	Informations http.HandlerFunc
}

type Server struct {
	Controllers
}

func New(c Controllers) *Server {
	return &Server{c}
}

func (s *Server) Run(p string) {
	fmt.Printf("start server on port %s\n", p)

	httpServer := &http.Server{
		Addr: fmt.Sprintf(":%s", p),
	}

	loadRoutes(s.Controllers)

	httpServer.ListenAndServe()
}
