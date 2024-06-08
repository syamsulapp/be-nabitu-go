package user

import (
	"be-nabitu-go/helpers"

	"github.com/gofiber/fiber/v2"
	jtoken "github.com/golang-jwt/jwt/v4"
)

func GetUser(c *fiber.Ctx) error {
	user := c.Locals("user").(*jtoken.Token)
	claim := user.Claims.(jtoken.MapClaims)
	mapUser := fiber.Map{
		"email":      claim["username"].(string),
		"username":   claim["email"].(string),
		"created_at": claim["created_at"].(string),
		"updated_at": claim["updated_at"].(string),
	}
	return c.JSON(helpers.SuccessResponse(c, "Berhasil Get Profile", mapUser))
}
