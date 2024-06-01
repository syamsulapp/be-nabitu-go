package repository

import "github.com/gofiber/fiber/v2"

func ResultGetRegisterRepository(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "halo ini bagian register",
	})
}
