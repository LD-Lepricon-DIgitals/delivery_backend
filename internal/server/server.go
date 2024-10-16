package server

import (
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/handlers"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/middleware"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
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
	s.srv.Use(recover.New())
	s.srv.Use(cors.New(cors.Config{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{"Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,X-CSRF-Token,Authorization"},
		AllowMethods: []string{"GET,POST,PATCH,DELETE"},
	}))
	log.Println("Starting server... Let`s Go :)")
	s.InitRoutes()
	log.Fatal(s.srv.Listen(s.cfg.HostAddr + ":" + s.cfg.HostPort))
}

func (s *Server) InitRoutes() {

	auth := s.srv.Group("/auth")
	auth.Post("/login", s.h.LoginUser)
	auth.Post("/register", s.h.RegisterUser)

	api := s.srv.Group("/api")
	user := api.Group("/user", s.mdl.AuthMiddleware) // TODO: add middleware
	user.Post("/change", s.h.ChangeUserCredentials)
	user.Post("/change_password", s.h.ChangeUserPassword)
	user.Post("/delete", s.h.DeleteUser)
	user.Post("/logout", s.h.LogoutUser)
	dishes := api.Group("/dishes")
	dishes.Get("/", s.h.GetDishes)
	dishes.Get("/by_id/:dish_id", s.h.GetDishById)
	dishes.Get("/by_category", s.h.GetDishesByCategory)
	dishes.Get("/search", s.h.SearchByName)
	secureDishes := dishes.Group("/") // TODO : add middleware
	secureDishes.Post("/add", s.h.AddDish)
	secureDishes.Delete("/delete/:id", s.h.DeleteDish)
	secureDishes.Put("/update", s.h.ChangeDish)
}

func (s *Server) Stop() {
	s.srv.Shutdown()
}
