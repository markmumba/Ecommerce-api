package main

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/markmumba/fiber-api/database"
	"github.com/markmumba/fiber-api/router"
)

func main() {
	database.ConnectDB()

	app := fiber.New()

	router.SetupRouter(app)

	err := app.Listen(":4000")

	if err != nil {
		log.Println(err.Error())
	}
}
