package database

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/markmumba/fiber-api/config"
	"github.com/markmumba/fiber-api/internals/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type DatabaseInstance struct {
	DB *gorm.DB
}

var Database DatabaseInstance

func ConnectDB() {

	var err error
	p := config.Config("DB_PORT")
	port, err := strconv.ParseUint(p, 10, 32)

	if err != nil {
		log.Println("wrongly configure port ")

	}

	database_string := fmt.Sprintf("host=%s port=%d user=%s dbname=%s sslmode= disable", config.Config("DB_HOST"), port, config.Config("DB_USER"), config.Config("DB_NAME"))
	db,err := gorm.Open(postgres.Open(database_string))
	if err != nil {
		log.Fatal("Failed to connect to the database! \n ", err.Error())
		os.Exit(2)
	}
	db.AutoMigrate(&models.User{},&models.Product{},&models.Order{})

	Database = DatabaseInstance{DB: db}

}
