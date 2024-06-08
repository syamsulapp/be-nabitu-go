package routes

import (
	auth "be-nabitu-go/handlers/auth"
	profile "be-nabitu-go/handlers/profile"
	user "be-nabitu-go/handlers/user"
	"be-nabitu-go/middlewares"
	"be-nabitu-go/pkg"

	"github.com/gofiber/fiber/v2"
)

var InitRouteAPI = func(app *fiber.App) {
	//init api v1
	routeGroup := app.Group("/api/v1")

	//init auth user api
	authUserRoute := routeGroup.Group("/auth/user")
	authUserRoute.Post("/register", auth.ResultHandleAuthRegister)
	authUserRoute.Post("/login", auth.ResultHandleAuthLogin)

	//init auth admin api
	// authAdminRoute := routeGroup.Group("/auth/admin")
	// authAdminRoute.Get("/login")

	//init profile route
	profileRoute := routeGroup.Group("/profile")
	profileRoute.Get("/", profile.ResultHandlerGetAllProfile)
	profileRoute.Post("/", profile.ResultHandlerCreateProfile)
	profileRoute.Get("/:id", profile.ResultHandlerGetProfileByID)

	//init api get user for session
	jwt := middlewares.InitMiddlewareAuth(pkg.Secret)
	routeGroup.Get("/user", jwt, user.GetUser)

}
