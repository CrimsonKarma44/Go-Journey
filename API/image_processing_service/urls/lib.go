package urls

import (
	"IPS/auth/handlers"
	"IPS/auth/services"
	handlers2 "IPS/img/handlers"
	services2 "IPS/img/services"
	"IPS/middleware"
	"fmt"
	"net/http"

	"gorm.io/gorm"
)

func Routes(DB *gorm.DB, SecretKeyAccess []byte, SecretKeyRefresh []byte, path string) {
	fmt.Println("Server Initializing...")

	store := make(map[string]string)

	// services
	authService := services.AuthService{DB: DB}.New()
	imgService := services2.ImageService{}.New(DB)

	// authHandlers
	authHandlers := handlers.AuthHandler{AuthService: authService, JwtSecretKeyAccess: SecretKeyAccess, JwtSecretKeyRefresh: SecretKeyRefresh, RefreshStore: store}

	// imgHandlers
	imgHandler := handlers2.ImageHandler{}.New(path, imgService)

	// AuthMiddleware
	AuthMiddleware := middleware.AuthMiddleware{JwtSecretKeyAccess: SecretKeyAccess, JwtSecretKeyRefresh: SecretKeyRefresh, RefreshStore: store}

	// routes
	mux := http.NewServeMux()
	// Auth routes
	mux.Handle("/auth/register", http.HandlerFunc(authHandlers.RegisterHandler))
	mux.Handle("/auth/login", http.HandlerFunc(authHandlers.LoginHandler))
	mux.Handle("/auth/logout", AuthMiddleware.ProtectMiddleware(http.HandlerFunc(authHandlers.LogoutHandler)))
	mux.Handle("/auth/renew-token", AuthMiddleware.RenewTokenMiddleware(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {})))

	// Image processing routes
	mux.Handle("/image/upload", AuthMiddleware.ProtectMiddleware(imgHandler.UploadHandler))
	mux.Handle("/image/{id}/transform", AuthMiddleware.ProtectMiddleware(imgHandler.TransformHandler))
	mux.Handle("/image/{id}/{format}", AuthMiddleware.ProtectMiddleware(imgHandler.RetrieveHandler))
	mux.Handle("/image", AuthMiddleware.ProtectMiddleware(imgHandler.RetrievePageHandler))

	// static Dir
	mux.Handle("/static/", AuthMiddleware.ProtectMiddleware(http.StripPrefix("/static/", http.FileServer(http.Dir(path[1:]))).ServeHTTP))

	http.ListenAndServe(":8080", mux)
}
