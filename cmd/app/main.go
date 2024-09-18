package main

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/server"
)

func main() {
	// TODO: init config
	// TODO: init logger
	// TODO: run server
	srv := server.NewServer()
	srv.Run("8080")
}
