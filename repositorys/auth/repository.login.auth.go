package repository

import "github.com/gofiber/fiber/v2"

func LoginRepository(c *fiber.Ctx) error {
	return c.JSON(fiber.Map{
		"message": "welcome api login",
	})
}
