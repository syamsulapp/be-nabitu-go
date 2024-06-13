package main

import (
	"be-nabitu-go/configs"
	"be-nabitu-go/database"
	"be-nabitu-go/routes"
	"be-nabitu-go/schemas"
	"log"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/qinains/fastergoding"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	//config .env
	configs.InitConfigEnv()
	// configs.CorsConfig()

	if os.Getenv("GO_ENV") != "stagging" && os.Getenv("GO_ENV") != "production" {
		//hot reload
		fastergoding.Run()
	}
	// //config cors
	// configs.CorsConfig()
	//config timezone
	if err := configs.InitConfigTimeZone(os.Getenv("GO_TIMEZONE")); err != nil {
		log.Fatal(err)
	}
	t := configs.GetConfigTimeZone(time.Now())
	log.Println(t.Format(os.Getenv("GO_TIMEZONE_FORMAT")))
	//config db
	database.InitConnectionDBMysql(configs.InitConfigDbMysql())
	// setup mux router
	app := SetupRouter()
	app.Listen(os.Getenv("GO_PORT"))

}

func indexRoute(c *fiber.Ctx) error {
	return c.JSON(schemas.SchemaIndexProfile{Message: "Welcome API Nabitu 1.0"})
}

func SetupRouter() *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	app.Get("/", indexRoute)
	routes.InitRouteAPI(app)
	return app
}
