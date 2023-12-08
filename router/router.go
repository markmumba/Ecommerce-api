package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/markmumba/ecommerceapp/internals/routes"
)

func SetupRouter(app *fiber.App) {
	api := app.Group("/api", logger.New())

	routes.SetupUserRoutes(api)
	routes.SetupProductRoutes(api)
	routes.SetupOrderRoutes(api)
}
