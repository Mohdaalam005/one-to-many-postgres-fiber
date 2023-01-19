package database

import (
	"github.com/mohdaalam005/one-to-many/helpers"
	"github.com/mohdaalam005/one-to-many/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func Connect() {
	url := os.Getenv("url")
	var err error
	DB, err = gorm.Open(postgres.Open(url), &gorm.Config{})
	helpers.CheckNilErr(err)
	DB.AutoMigrate(models.Student{}, models.Department{})
}
