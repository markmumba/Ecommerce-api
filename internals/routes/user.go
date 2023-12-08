package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/markmumba/ecommerceapp/internals/handlers"
)

func SetupUserRoutes(router fiber.Router) {

	users := router.Group("/users")

	users.Get("/", handlers.GetUsers)
	users.Post("/create", handlers.CreateUser)
	users.Get("/:id", handlers.GetUser)
	users.Put("/:id", handlers.UpdateUser)
	users.Delete("/:id", handlers.DeleteUser)

}
