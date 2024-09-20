package main

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/server"
)

func main() {

	cfg := config.NewConfig() //config initialization

	srv := server.NewServer(cfg) //server initialization

	srv.Run() //server start

}
