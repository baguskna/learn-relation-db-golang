package config

import (
	"learn-relation-db-go/models"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func StartDB() {
	dsn := "root:@tcp(127.0.0.1:3306)/learn-relation-db-go?charset=utf8mb4&parseTime=True&loc=Local"
	dbRes, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("DB connection error")
	}

	db = dbRes

	// auto migrate
	db.AutoMigrate(&models.User{}, &models.Order{})
}

func GetDB() *gorm.DB {
	return db
}
