package middleware

import (
	"IPS/auth"
	"IPS/auth/model"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthMiddleware struct {
	JwtSecretKeyAccess  []byte
	JwtSecretKeyRefresh []byte

	RefreshStore map[string]string
}

func (a *AuthMiddleware) RenewTokenMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("refresh_token")
		if err != nil {
			http.Error(w, "no refresh token", http.StatusUnauthorized)
			return
		}

		refreshToken := cookie.Value
		claims := &model.Claims{}
		token, err := jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
			return a.JwtSecretKeyRefresh, nil
		})
		if err != nil || !token.Valid {
			http.Error(w, "invalid refresh token", http.StatusUnauthorized)
			return
		}

		// Check refresh store
		if stored, ok := a.RefreshStore[claims.Email]; !ok || stored != refreshToken {
			http.Error(w, "refresh token revoked", http.StatusUnauthorized)
			return
		}

		// Generate new access token
		access, newRefresh, err := auth.GenerateToken(claims.Email, a.JwtSecretKeyAccess, a.JwtSecretKeyRefresh, a.RefreshStore)
		if err != nil {
			http.Error(w, "could not generate new token", http.StatusInternalServerError)
			return
		}

		// Rotate refresh token
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    newRefresh,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   false,
			Path:     "/",
		})

		json.NewEncoder(w).Encode(map[string]string{
			"access_token": access,
		})

		next.ServeHTTP(w, r)
	})
}

func (a *AuthMiddleware) ProtectMiddleware(next http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			http.Error(w, "Missing token", http.StatusUnauthorized)
			return
		}
		// Remove "Bearer "
		if len(tokenString) > 7 && tokenString[:7] == "Bearer " {
			tokenString = tokenString[7:]
		}

		// Validate the existing token
		_, err := auth.ValidateJWT(tokenString, a.JwtSecretKeyAccess)
		if err != nil {
			http.Error(w, "Invalid token", http.StatusUnauthorized)
			return
		}

		// Call the next handler
		next.ServeHTTP(w, r)
	})
}
