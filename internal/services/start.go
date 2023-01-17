package services

import (
	"log"

	"github.com/BoiseITGuru/ArrRequests/internal/config"
	"github.com/BoiseITGuru/ArrRequests/internal/models"
	"github.com/BoiseITGuru/ArrRequests/sdk/tmdb"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var TMDB *tmdb.Client

func Start() {
	startDB()
	startTMDB()
}

func startDB() {
	instance, err := gorm.Open(sqlite.Open("ArrRequests.db"), &gorm.Config{})
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to Database!")

	err = instance.AutoMigrate(
		&models.User{},
	)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Database Migration Completed!")

	DB = instance
}

func startTMDB() error {
	client, err := tmdb.NewClient(nil, config.AppConfig.TmdbV4Key)
	if err != nil {
		return err
	}

	TMDB = client

	return nil
}
