package main

import (
	"nearby/health"
	"nearby/server"
)

func main() {
	c := server.Controllers{
		Health: health.Controller,
	}

	s := server.New(c)
	s.Run("8080")
}
