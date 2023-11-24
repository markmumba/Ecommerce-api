package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/markmumba/fiber-api/internals/routes"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	routes.SetupUserRoutes(api)
}
