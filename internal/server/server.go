package server

import (
	"github.com/gofiber/fiber/v3/middleware/adaptor"
	httpSwagger "github.com/swaggo/http-swagger"
	"log"

	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/config"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/handlers"
	"github.com/LD-Lepricon-DIgitals/delivery_backend/internal/middleware"
	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
	"github.com/gofiber/fiber/v3/middleware/recover"
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
		AppName:         "Delivery app ver 1.0",
		ErrorHandler:    handlers.CustomError,
		StructValidator: &structValidator{validate: validator.New()},
	})

	// Apply Middlewares
	s.applyMiddlewares()

	log.Println("Starting server... Let`s Go :)")
	s.InitRoutes()
	log.Fatal(s.srv.Listen(s.cfg.HostAddr + ":" + s.cfg.HostPort))
}

func (s *Server) applyMiddlewares() {
	s.srv.Use(logger.New())
	s.srv.Use(recover.New())
	s.srv.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowHeaders:     []string{"Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,X-CSRF-Token,Authorization"},
		AllowMethods:     []string{"GET,POST,PATCH,DELETE"},
		AllowCredentials: true,
	}))
}

func (s *Server) InitRoutes() {
	auth := s.srv.Group("/auth")
	s.initAuthRoutes(auth)

	api := s.srv.Group("/api")
	api.Get("/swagger/", adaptor.HTTPHandlerFunc(httpSwagger.WrapHandler))

	user := api.Group("/user", s.mdl.AuthMiddleware)
	s.initUserRoutes(user)

	dishes := api.Group("/dishes")
	s.initDishRoutes(dishes)

	secureAdmin := api.Group("/secure", s.mdl.AdminAuthMiddleware)
	s.initAdminRoutes(secureAdmin)

	secureUserOrder := api.Group("/user", s.mdl.AuthMiddleware)
	s.initUserOrderRoutes(secureUserOrder)
}

func (s *Server) initAuthRoutes(group fiber.Router) {
	group.Post("/login", s.h.LoginUser)
	group.Post("/register", s.h.RegisterUser)
}

func (s *Server) initUserRoutes(group fiber.Router) {
	group.Patch("/change", s.h.ChangeUserCredentials)
	group.Patch("/change_password", s.h.ChangeUserPassword)
	group.Delete("/delete", s.h.DeleteUser)
	group.Post("/logout", s.h.LogoutUser)
	group.Get("/info", s.h.GetUserInfo)
	group.Patch("/photo", s.h.UpdatePhoto)
}

func (s *Server) initDishRoutes(group fiber.Router) {
	group.Get("/", s.h.GetDishes)
	group.Get("/by_id/:dish_id", s.h.GetDishById)
	group.Post("/by_category", s.h.GetDishesByCategory)
	group.Get("/search/:name", s.h.SearchByName)
}

func (s *Server) initAdminRoutes(group fiber.Router) {
	group.Post("/add", s.h.AddDish)
	group.Delete("/delete/:id", s.h.DeleteDish)
	group.Put("/update", s.h.ChangeDish)
}

func (s *Server) initUserOrderRoutes(group fiber.Router) {
	group.Post("/create_order", s.h.CreateOrderHandler)
	group.Delete("/delete_order/:orderId", s.h.DeleteOrder)
	group.Get("/orders", s.h.GetUserOrders)
}

func (s *Server) Stop() {
	s.srv.Shutdown()
}
