package ecommerce

import (
	"ecommerce/database"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// Init function initializes the database using the GORM ORM
func Init() *gorm.DB {
	// Connect to the database
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// Migrate the schema
	err = db.AutoMigrate(&database.User{}, &database.Product{}, &database.Cart{}, &database.Order{}, &database.OrderList{})
	if err != nil {
		panic("failed to connect database")
	}
	return db
}
