package handlers

import (
	"errors"

	fiber "github.com/gofiber/fiber/v2"
	"github.com/markmumba/ecommerceapp/database"
	"github.com/markmumba/ecommerceapp/internals/models"
)

type ProductSerializer struct {
	ID           uint   `json:"id"`
	Name         string `json:"name"`
	Image        string `json:"image"`
	Description  string `json:"description"`
	SerialNumber string `json:"serial_number"`
}

func CreateSerialProduct(product models.Product) ProductSerializer {
	return ProductSerializer{ID: product.ID, Name: product.Name, Image: product.Image, Description: product.Description, SerialNumber: product.SerialNumber}
}

func AddProduct(c *fiber.Ctx) error {
	var product models.Product

	err := c.BodyParser(&product)
	if err != nil {
		c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&product)
	responseProduct := CreateSerialProduct(product)

	return c.Status(200).JSON(responseProduct)

}

func GetProducts(c *fiber.Ctx) error {

	products := []models.Product{}

	database.Database.DB.Find(&products)

	var newProduct []ProductSerializer
	for _, product := range products {

		responseProduct := CreateSerialProduct(product)
		newProduct = append(newProduct, responseProduct)

	}

	return c.Status(200).JSON(newProduct)

}

func FindProduct(id int, product *models.Product) error {

	database.Database.DB.Find(&product, "id=?", id)
	if product.ID == 0 {
		return errors.New("There are no results man ")
	}

	return nil

}

func GetProduct(c *fiber.Ctx) error {
	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(500).JSON("no id has been given")
	}

	err = FindProduct(id, &product)

	responseProduct := CreateSerialProduct(product)

	return c.Status(200).JSON(responseProduct)

}

func UpdateProduct(c *fiber.Ctx) error {

	var product models.Product
	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(500).JSON("No id was read ")
	}

	err = FindProduct(id, &product)

	type UpdateProduct struct {
		Name         string `json:"name"`
		SerialNumber string `json:"serial_number"`
	}
	var updateproduct UpdateProduct

	if err := c.BodyParser(&updateproduct); err != nil {
		c.Status(500).JSON(err.Error())
	}

	product.Name = updateproduct.Name
	product.SerialNumber = updateproduct.SerialNumber

	responseProduct := CreateSerialProduct(product)

	return c.Status(200).JSON(responseProduct)
}

func DeleteProduct(c *fiber.Ctx) error {
	var product models.Product

	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(500).JSON("no id has been given")
	}
	err = FindProduct(id, &product)

	database.Database.DB.Delete(&product)

	allProducts := GetProducts(c)

	return c.Status(200).JSON(allProducts)
}
