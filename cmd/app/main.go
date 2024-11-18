package main

import (
	_ "github.com/LD-Lepricon-DIgitals/delivery_backend/docs"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/handlers"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/middleware"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/server"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/db"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/pkg/service"
	"log"
)

// @title Delivery Backend API
// @version 1.0
// @description API documentation for the Delivery Backend
// @contact.name API Support
// @contact.email support@example.com
// @host localhost:1317
// @BasePath /
func main() {

	cfg := config.NewConfig()

	database, err := db.NewDBConn(cfg)

	if err != nil {
		log.Fatal(err)
	}

	log.Println("Connected to database")

	repo := db.NewRepository(database) //repo will take db as an argument

	services := service.NewService(repo, cfg)

	middlw := middleware.NewMiddleware(services)
	handler := handlers.NewHandlers(services)

	srv := server.NewServer(cfg, handler, middlw)

	srv.Run() //server start

	defer srv.Stop()

	///google cloud
	//pagination
	//
}
