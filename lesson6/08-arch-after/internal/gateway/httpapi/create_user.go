package httpapi

import (
	"github.com/gofiber/fiber/v2"
)

func (s *Server) CreateUser(c *fiber.Ctx) error {
	name := c.Params("name")
	u, err := s.uc.CreateUser(c.Context(), name)
	if err != nil {
		return err
	}
	return c.JSON(u)
}
