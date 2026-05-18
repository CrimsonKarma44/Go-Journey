package middleware

import (
	"ecommerce/utils"
	"fmt"
	"net/http"
)

func RenewTokenMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Validate the existing token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			//w.WriteHeader(http.StatusForbidden)
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Generate a new token with an updated expiration time
		newToken, err := utils.GenerateToken(claims.Username)
		if err != nil {
			http.Error(w, "Failed to generate new token", http.StatusInternalServerError)
			return
		}

		// Attach the new token to the response header
		w.Header().Set("Authorization", newToken)

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
func ProtectMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}

		// Validate the existing token
		claims, err := utils.ValidateJWT(tokenString)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}
		fmt.Println(claims.Username)
		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
