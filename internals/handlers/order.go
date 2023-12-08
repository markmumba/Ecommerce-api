package handlers

import (
	"errors"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/markmumba/ecommerceapp/database"
	"github.com/markmumba/ecommerceapp/internals/models"
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

func GetOrders(c *fiber.Ctx) error {

	orders := []models.Order{}

	database.Database.DB.Find(&orders)

	var orderlist []OrderSerializer

	for _, order := range orders {
		var user models.User
		var product models.Product

		database.Database.DB.Find(&user, "id=?", order.UserRefer)
		database.Database.DB.Find(&product, "id=?", order.ProductRefer)
		responseOrder := CreateSerialOrder(order, CreateSerialUser(user), CreateSerialProduct(product))

		orderlist = append(orderlist, responseOrder)

	}

	return c.Status(200).JSON(orderlist)

}

func FindOrder(id int, order *models.Order) error {
	database.Database.DB.Find(&order, "id=?", id)
	if order.ID == 0 {
		return errors.New("no order of that id")
	}

	return nil
}

func GetOrder(c *fiber.Ctx) error {
	var order models.Order

	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(400).JSON("no id given")
	}
	err = FindOrder(id, &order)

	var user models.User
	var product models.Product

	database.Database.DB.First(&user, order.UserRefer)
	database.Database.DB.First(&product, order.ProductRefer)
	responseUser := CreateSerialUser(user)
	responseProduct := CreateSerialProduct(product)
	responseOrder := CreateSerialOrder(order, responseUser, responseProduct)

	return c.Status(200).JSON(responseOrder)
}
