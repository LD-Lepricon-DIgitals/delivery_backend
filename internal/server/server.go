package server

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/handlers"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"log"
)

type Server struct {
	srv *fiber.App
	cfg *config.Config
	h   *handlers.Handlers
	mdl *middleware.Middleware
}

func NewServer(config *config.Config, handlers *handlers.Handlers, midd *middleware.Middleware) *Server {
	return &Server{
		cfg: config,
		h:   handlers,
		mdl: midd,
	}
}

func (s *Server) Run() {
	s.srv = fiber.New(fiber.Config{
		AppName:      "Delivery app ver 1.0",
		ErrorHandler: handlers.CustomError,
	})
	s.srv.Use(logger.New())
	s.srv.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,X-CSRF-Token,Authorization"},
		AllowMethods: []string{"GET,POST,PATCH,DELETE"},
	}))
	log.Println("Starting server... Let`s Go :)")
	s.InitRoutes()
	log.Fatal(s.srv.Listen(":" + s.cfg.HostPort))
}

func (s *Server) InitRoutes() {

	auth := s.srv.Group("/auth")
	auth.Post("/login", s.h.LoginUser)
	auth.Post("/register", s.h.RegUser)

	api := s.srv.Group("/api")
	user := api.Group("/user", s.mdl.AuthMiddleware) // TODO: add middleware
	user.Get("/profile::id", s.h.GetUserInfo)        //TODO: get method + query params
	user.Post("/add_info", s.h.AddUserInfo)
	user.Post("/add_address", s.h.AddUserAddress)
	user.Put("/change_city", s.h.ChangeUserCity)
	user.Put("/change_email", s.h.ChangeUserEmail)
	user.Put("/change_password", s.h.ChangeUserPassword)
	user.Put("/change_login", s.h.ChangeUserLogin)
	user.Delete("/delete", s.h.DeleteUser)
	user.Put("/change_phone", s.h.ChangeUserPhone)

}

func (s *Server) Stop() {
	s.srv.Shutdown()
}
