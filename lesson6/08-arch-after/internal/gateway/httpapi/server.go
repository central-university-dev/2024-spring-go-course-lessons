package httpapi

import (
	"github.com/gofiber/fiber/v2"

	"github.com/klakovsky/users/internal/usecase"
)

type Server struct {
	uc *usecase.Usecase
}

func NewServer(us *usecase.Usecase) *Server {
	return &Server{uc: us}
}

func (s *Server) Run() {
	app := fiber.New()
	app.Get("/users/:user_id", s.GetUser)
	app.Post("/users/:name", s.CreateUser)
	app.Listen(":3000")
}
