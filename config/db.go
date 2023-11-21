package config

import (
	"github.com/mokshmalik5757/Go-Rest-API/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)
var DB *gorm.DB

func Connect() {
	db, err := gorm.Open(postgres.Open("postgres://postgres:postgres@localhost:5433/postgres"), &gorm.Config{})

	if err!= nil {
		panic(err)
	}
	db.AutoMigrate(&models.User{})
	DB = db
}