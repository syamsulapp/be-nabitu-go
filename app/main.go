package main

import (
	"be-nabitu-go/configs"
	"be-nabitu-go/database"
	"be-nabitu-go/routes"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
)

func main() {
	//hot reload
	fastergoding.Run()
	//config .env
	configs.InitConfigEnv()
	//config db
	database.InitConnectionDBMysql(configs.InitConfigDbMysql())
	// setup mux router
	app := SetupRouter()
	app.Listen(os.Getenv("GO_PORT"))

}

func SetupRouter() *fiber.App {
	app := fiber.New()
	db := database.GetDBMysql()
	routes.InitRouteProfile(app, db)
	return app
}
