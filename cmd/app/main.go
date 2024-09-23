package main

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/handlers"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/server"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"
	"log"
)

func main() {

	cfg := config.NewConfig()

	database, err := db.NewDBConn(cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	repo := db.NewRepository(database) //repo will take db as an argument

	services := service.NewService(repo)
	handler := handlers.NewHandlers(services)

	srv := server.NewServer(cfg, handler)

	srv.Run() //server start

	defer srv.Stop()

	///google cloud
	//pagination
	//
}
