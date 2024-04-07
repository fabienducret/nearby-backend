package server

import (
	"net/http"
)

func loadRoutes(c Controllers) {
	http.HandleFunc("GET /health", c.Health)
	http.HandleFunc("GET /informations/{city}", c.Informations)
}
