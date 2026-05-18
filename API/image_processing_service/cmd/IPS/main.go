package main

import (
	"IPS/auth/model"
	"IPS/urls"
	"IPS/utils"

	"fmt"
)

const (
	LocalStorePath = "/static"
)

func main() {
	// Initialize local storage
	err := utils.InitLocalStore(LocalStorePath)
	if err != nil {
		panic("Failed to initialize local storage!")
	}
	fmt.Println("Local storage initialized successfully.")

	// Initialize database
	data := model.Data_values{}.Load()
	db, err := utils.InitDB(data)
	if err != nil {
		panic("Failed to connect to database!")
	}
	fmt.Println("Database connected successfully.")

	// Migrate the schema
	if err := utils.MigrateDB(db); err != nil {
		panic("Failed to migrate database schema!")
	}
	fmt.Println("Database schema migrated successfully.")

	// Initialize routes
	urls.Routes(db, data.SECRET_KEY_Access, data.SECRET_KEY_Refresh, LocalStorePath)

	fmt.Println("Server started on :8080")
}
