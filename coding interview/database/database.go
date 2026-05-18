package database

import (
	"coding_interview/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
)

var DB *gorm.DB

func Connect() {
	dsn := "user:password@tcp(127.0.0.1:3306)/library?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	log.Println("Connected to database")

	// Auto migrate models
	db.AutoMigrate(&models.User{}, &models.Book{}, &models.Borrow{})
	DB = db
}
