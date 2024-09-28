package handler

import "github.com/gofiber/fiber/v3"

func HealthCheck(c fiber.Ctx) error {

	c.Status(fiber.StatusOK).JSON(fiber.Map{
		"message": "Wake up, Neo...",
	})

	return nil
}
