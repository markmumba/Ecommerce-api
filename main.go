package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/markmumba/ecommerceapp/database"
	"github.com/markmumba/ecommerceapp/router"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://localhost:5173",
		AllowHeaders: "Origin,Content-Type,Accept",
	}))

	router.SetupRouter(app)

	err := app.Listen(":4000")

	if err != nil {
		log.Println(err.Error())
	}
}
