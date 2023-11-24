package handlers

import (
	"errors"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/markmumba/fiber-api/database"
	"github.com/markmumba/fiber-api/internals/models"
)

type UserSerializer struct {
	ID        uint   `json:"id" `
	FirstName string `json:"first_name"`
	LastName  string ` json:"last_name"`
}

func ErrorHelper(c *fiber.Ctx, message string) error {
	return c.Status(400).JSON(message)

}

func CreateSerialUser(userModel models.User) UserSerializer {
	return UserSerializer{ID: userModel.ID, FirstName: userModel.FirstName, LastName: userModel.LastName}
}

func CreateUser(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)
	fmt.Println(user)
	if err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.Database.DB.Create(&user)
	responseUser := CreateSerialUser(user)

	return c.Status(200).JSON(responseUser)

}
func GetUsers(c *fiber.Ctx) error {

	users := []models.User{}
	database.Database.DB.Find(&users)
	responseUsers := []UserSerializer{}

	for _, user := range users {
		responseUser := CreateSerialUser(user)
		responseUsers = append(responseUsers, responseUser)
	}
	return c.Status(200).JSON(responseUsers)
}

func FindUser(id int, user *models.User) error {
	database.Database.DB.Find(&user, "id= ?", id)
	if user.ID == 0 {
		return errors.New("user does not exist")
	}
	return nil
}
func GetUser(c *fiber.Ctx) error {
	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return ErrorHelper(c, "Please ensure that id is an integer")
	}
	err = FindUser(id, &user)
	if err != nil {
		return ErrorHelper(c, "Not available")
	}
	responseUser := CreateSerialUser(user)
	return c.Status(200).JSON(responseUser)

}

func UpdateUser(c *fiber.Ctx) error {

	id, err := c.ParamsInt("id")
	var user models.User
	if err != nil {
		return ErrorHelper(c, "Please ensure that id is an integer")
	}
	err = FindUser(id, &user)
	if err != nil {
		return ErrorHelper(c, "Not available")
	}

	type UpdateUser struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
	}
	var updateuser UpdateUser

	if err = c.BodyParser(&updateuser); err != nil {
		c.Status(500).JSON(err.Error())
	}

	user.FirstName = updateuser.FirstName
	user.LastName = updateuser.LastName

	database.Database.DB.Save(&user)

	responseUser := CreateSerialUser(user)

	return c.Status(200).JSON(responseUser)
}

func DeleteUser(c *fiber.Ctx) error {
	var user models.User

	id, err := c.ParamsInt("id")
	if err != nil {
		c.Status(400).JSON("the id is not correct ")
	}
	err = FindUser(id, &user)
	database.Database.DB.Delete(&user)
	answer := GetUsers(c)
	return answer

}
