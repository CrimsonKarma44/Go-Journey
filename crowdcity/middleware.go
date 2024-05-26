package main

import (
	"net/http"
)

// AuthMiddleware is a middleware for protecting routes.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil {
			if err == http.ErrNoCookie {
				http.Error(w, "No token found", http.StatusUnauthorized)
				return
			}
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		tokenString := cookie.Value
		_, err = ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
