package server

import (
	"github.com/gofiber/fiber/v2"
	"log"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run(port string) {
	srv := fiber.New()
	log.Fatal(srv.Listen(":" + port))
	return
}
