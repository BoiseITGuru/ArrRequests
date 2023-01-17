package database

import (
	"log"

	"github.com/BoiseITGuru/ArrRequests/internal/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Instance *gorm.DB
var dbError error

func Connect() {
	Instance, dbError = gorm.Open(sqlite.Open("ArrRequests.db"), &gorm.Config{})
	if dbError != nil {
		log.Println(dbError)
		panic("Cannot connect to DB")
	}
	log.Println("Connected to Database!")
}

func Migrate() {
	err := Instance.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Migration Completed!")
}
