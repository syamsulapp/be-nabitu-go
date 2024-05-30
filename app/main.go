package main

import (
	"be-nabitu-go/configs"
	"be-nabitu-go/database"
	"be-nabitu-go/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/qinains/fastergoding"
)

func main() {
	//config .env
	configs.InitConfigEnv()

	if os.Getenv("GO_ENV") != "stagging" && os.Getenv("GO_ENV") != "production" {
		//hot reload
		fastergoding.Run()
	}
	//config db
	database.InitConnectionDBMysql(configs.InitConfigDbMysql())
	// setup mux router
	app := SetupRouter()
	app.Listen(os.Getenv("GO_PORT"))

}

func SetupRouter() *fiber.App {
	app := fiber.New()
	db := database.GetDBMysql()

	app.Use(cors.New(cors.Config{
		AllowHeaders:     "Origin,Content-Type,Accept,Content-Length,Accept-Language,Accept-Encoding,Connection,Access-Control-Allow-Origin,XMLHttpRequest",
		AllowOriginsFunc: func(origin string) bool { return true },
		AllowCredentials: true,
		AllowMethods:     "GET, POST, PUT, DELETE, PATCH, OPTIONS",
	}))
	routes.InitRouteProfile(app, db)
	return app
}
