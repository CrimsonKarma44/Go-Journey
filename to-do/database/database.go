package database

import (
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"to-do/models"
)

func DNS() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_NAME"),
	)
	return dsn
}

func GormInit() (*gorm.DB, error) {
	DB, err := gorm.Open(mysql.Open(DNS()), &gorm.Config{})
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
		return nil, err
	}

	log.Println("Database connection established")
	return DB, nil
}
func Migration(db *gorm.DB) error {
	err := db.AutoMigrate(&models.User{}, &models.List{})
	if err != nil {
		log.Fatalf("Error auto migrate lists: %v", err)
		return err
	}
	log.Println("Migration complete")
	return nil
}
