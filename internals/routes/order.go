package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markmumba/ecommerceapp/internals/handlers"
)

func SetupOrderRoutes(router fiber.Router) {
	orders := router.Group("/orders")

	orders.Get("/", handlers.GetOrders)
	orders.Post("/create", handlers.CreateOrder)
	orders.Get("/:id", handlers.GetOrder)
}
