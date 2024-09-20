package server

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/handlers"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server struct {
	srv *fiber.App
	cfg *config.Config
	h   *handlers.Handlers
}

func NewServer(config *config.Config, handlers *handlers.Handlers) *Server {
	return &Server{
		cfg: config,
		h:   handlers,
	}
}

func (s *Server) Run() {
	s.srv = fiber.New()
	log.Println("Starting server... Let`s Go :)")
	log.Fatal(s.srv.Listen(":" + s.cfg.HostPort))
	return
}

func (s *Server) InitRoutes() error {
	return nil
}

func (s *Server) Stop() {
	s.srv.Shutdown()
}
