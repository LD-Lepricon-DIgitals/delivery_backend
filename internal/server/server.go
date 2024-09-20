package server

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server struct {
	srv *fiber.App
	cfg *config.Config
}

func NewServer(config *config.Config) *Server {
	return &Server{
		cfg: config,
	}
}

func (s *Server) Run() {
	s.srv = fiber.New()
	log.Fatal(s.srv.Listen(":" + s.cfg.HostPort))
	return
}
