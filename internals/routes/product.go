package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markmumba/ecommerceapp/internals/handlers"
)

func SetupProductRoutes(router fiber.Router) {

	products := router.Group("/products")

	products.Get("/", handlers.GetProducts)
	products.Post("/create", handlers.AddProduct)
	products.Get("/:id", handlers.GetProduct)
	products.Put("/:id", handlers.UpdateProduct)
	products.Delete("/:id", handlers.DeleteProduct)

}
