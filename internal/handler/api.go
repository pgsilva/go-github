package handler

import (
	"github.com/gofiber/fiber/v3"
	"github.com/pgsilva/go-github/internal/usecase"
)

func SearchUsers(c fiber.Ctx) error {
	resp, err := usecase.SearchUsers(c.Query("q"))
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{
			"message": err.Error(),
		})
	}

	return c.Status(fiber.StatusOK).JSON(resp)
}
