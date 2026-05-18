package main

import (
	"coding_interview/database"
	"coding_interview/handlers"
	"coding_interview/repositories"
	"coding_interview/services"
	"github.com/gin-gonic/gin"
)

func main() {
	// Connect to the database
	database.Connect()

	// Initialize repositories
	userRepo := &repositories.UserRepository{DB: database.DB}
	bookRepo := &repositories.BookRepository{DB: database.DB}
	borrowRepo := &repositories.BorrowRepository{DB: database.DB}

	// Initialize services
	userService := &services.UserService{UserRepo: userRepo}
	bookService := &services.BookService{BookRepo: bookRepo}
	borrowService := &services.BorrowService{BorrowRepo: borrowRepo, BookRepo: bookRepo}

	// Initialize handlers
	userHandler := &handlers.UserHandler{UserService: userService}
	bookHandler := &handlers.BookHandler{BookService: bookService}
	borrowHandler := &handlers.BorrowHandler{BorrowService: borrowService}

	// Set up the router
	router := gin.Default()

	// User routes
	router.POST("/users", userHandler.EnrollUser)

	// Book routes
	router.GET("/books", bookHandler.ListBooks)
	router.GET("/books/:id", bookHandler.GetBookByID)
	router.GET("/books/filter", bookHandler.FilterBooks)

	// Borrow routes
	router.POST("/books/:id/borrow", borrowHandler.BorrowBook)

	// Start the server
	router.Run(":8080")
}
