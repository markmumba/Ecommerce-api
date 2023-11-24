package handlers

import (
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/markmumba/fiber-api/database"
	"github.com/markmumba/fiber-api/internals/models"
)

type OrderSerializer struct {
	ID        uint              `json:"id"`
	User      UserSerializer    `json:"user"`
	Product   ProductSerializer `json:"product"`
	CreatedAt time.Time
}

func CreateSerialOrder(order models.Order, user UserSerializer, product ProductSerializer) OrderSerializer {
	return OrderSerializer{ID: order.ID, User: user, Product: product, CreatedAt: order.CreatedAt}

}

func CreateOrder(c *fiber.Ctx) error {

	var order models.Order
	if err := c.BodyParser(&order); err != nil {
		c.Status(500).JSON(err.Error())
	}

	var user models.User
	if err := FindUser(order.UserRefer, &user); err != nil {
		c.Status(400).JSON(err.Error())
	}

	var product models.Product
	if err := FindProduct(order.ProductRefer, &product); err != nil {
		c.Status(400).JSON(err.Error())
	}
	database.Database.DB.Create(&order)

	responseUser := CreateSerialUser(user)
	responseProduct := CreateSerialProduct(product)
	responseOrder := CreateSerialOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)

}
