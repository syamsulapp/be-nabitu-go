package routes

import (
	handlers "be-nabitu-go/handlers/profile"
	repositorys "be-nabitu-go/repositorys/profile"
	"be-nabitu-go/schemas"
	services "be-nabitu-go/services/profile"

	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

func indexRoute(c *fiber.Ctx) error {
	return c.JSON(schemas.SchemaIndexProfile{Message: "Welcome API Nabitu 1.0"})
}

var InitRouteProfile = func(app *fiber.App, db *gorm.DB) {
	app.Get("/", indexRoute)

	//list profile
	resultProfileRepository := repositorys.NewRepositoryProfileResult(db)
	resultProfileService := services.NewServiceProfileResult(resultProfileRepository)
	resultProfileHandler := handlers.NewHandlerGetAllProfile(resultProfileService)

	routerGroup := app.Group("/api/v1")
	routerGroup.Get("/profile", resultProfileHandler.ResultProfileHandler)
}
