package urls

import (
	"ecommerce/handler"
	"ecommerce/middleware"
	"fmt"
	"gorm.io/gorm"
	"log"
	"net/http"
)

// Urls function manages the endpoint connections
func Urls(db *gorm.DB) {
	userHandler := handler.UserHandler{
		DB: db,
	}
	productHandler := handler.ProductHandler{
		DB: db,
	}
	cartHandler := handler.CartHandler{
		DB: db,
	}
	orderHandler := handler.OrderHandler{
		DB: db,
	}

	mux := http.NewServeMux()

	// User Handler Url
	mux.HandleFunc("/signup", userHandler.SignupHandler)
	mux.HandleFunc("/login", userHandler.LoginHandler)
	mux.Handle("/logout", middleware.ProtectMiddleware(userHandler.LogoutHandler))
	mux.Handle("/user/{id}", middleware.ProtectMiddleware(userHandler.ViewHandler))

	// Product Handler Url
	mux.Handle("/product/create", middleware.ProtectMiddleware(productHandler.CreateHandler))
	mux.Handle("/product/{id}", middleware.ProtectMiddleware(productHandler.ViewHandler))
	mux.Handle("/product/{id}/delete", middleware.ProtectMiddleware(productHandler.DeleteHandler))

	// Cart Handler Url
	mux.Handle("/product/{id}/to-cart", middleware.ProtectMiddleware(cartHandler.AddToCart))
	mux.Handle("/user/{id}/cart/", middleware.ProtectMiddleware(cartHandler.ViewCart))

	// Order and Payment url
	mux.Handle("/user/{id}/order/", middleware.ProtectMiddleware(orderHandler.Order))

	fmt.Println("Server Started...")
	log.Fatal(http.ListenAndServe(":8080", mux))
}
