package database

import (
	"MyGram/models"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	host     = "localhost"
	user     = "postgres"
	password = "indonesia00"
	dbPort   = "5432"
	dbname   = "mygram"
	db       *gorm.DB
	err      error
)

func StartDB() {
	config := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable", host, user, password, dbname, dbPort)
	dsn := config
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		log.Fatal("error connecting to database : ", err)
	}

	fmt.Println("connected")
	db.Debug().AutoMigrate(models.User{}, models.Photo{}, models.SocialMedia{}, models.Comment{})
}

func GetDB() *gorm.DB {
	return db
}
