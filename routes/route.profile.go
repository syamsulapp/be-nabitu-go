package routes

import (
	handlers "be-nabitu-go/handlers/profile"
	"be-nabitu-go/schemas"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func indexRoute(c *fiber.Ctx) error {
	return c.JSON(schemas.SchemaIndexProfile{Message: "Welcome API Nabitu 1.0"})
}

var InitRouteProfile = func(app *fiber.App, db *gorm.DB) {
	app.Get("/", indexRoute)

	routerGroup := app.Group("/api/v1")
	routerGroup.Get("/profile", handlers.ResultHandlerGetAllProfile)
}
