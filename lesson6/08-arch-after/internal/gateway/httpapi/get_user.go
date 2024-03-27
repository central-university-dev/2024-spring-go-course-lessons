package httpapi

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (s *Server) GetUser(c *fiber.Ctx) error {
	userID := c.Params("user_id")
	ID, err := strconv.Atoi(userID)
	if err != nil {
		return err
	}
	u, err := s.uc.GetUser(c.Context(), ID)
	if err != nil {
		return err
	}
	return c.JSON(u)
}
