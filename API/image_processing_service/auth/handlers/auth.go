package handlers

import (
	"IPS/auth"
	"IPS/auth/model"
	"IPS/auth/services"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type AuthHandler struct {
	// Add necessary fields, e.g., DB connection, auth service, etc.
	AuthService services.AuthService

	JwtSecretKeyAccess  []byte
	JwtSecretKeyRefresh []byte

	RefreshStore map[string]string
}

func (h *AuthHandler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	// Implement registration logic
	if r.Method == "POST" {
		var creds model.Creds
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		// Save to database (pseudo-code, replace with actual DB logic)
		res, err := h.AuthService.Register(&creds)
		if err != nil {
			http.Error(w, "could not save user", http.StatusInternalServerError)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write(res)
	}

}

func (h *AuthHandler) LoginHandler(w http.ResponseWriter, r *http.Request) {
	// Implement login logic
	var creds model.Creds
	if r.Method == "POST" {
		if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
			http.Error(w, "invalid request", http.StatusBadRequest)
			return
		}

		response, err := h.AuthService.Login(creds)
		if err != nil {
			http.Error(w, fmt.Sprintf("could not login: %s", err), http.StatusInternalServerError)
			return
		}

		// Generate tokens
		access, refresh, err := auth.GenerateToken(creds.Email, h.JwtSecretKeyAccess, h.JwtSecretKeyRefresh, h.RefreshStore)
		if err != nil {
			http.Error(w, "could not generate tokens", http.StatusInternalServerError)
			return
		}

		// Set refresh token as HttpOnly cookie
		http.SetCookie(w, &http.Cookie{
			Name:     "refresh_token",
			Value:    refresh,
			Expires:  time.Now().Add(24 * time.Hour),
			HttpOnly: true,
			Secure:   false, // true in production with HTTPS
			Path:     "/",
		})

		json.NewEncoder(w).Encode(map[string]interface{}{
			"access_token": access,
			"status":       http.StatusAccepted,
			"response":     json.RawMessage(response),
		})
	}
}

func (h *AuthHandler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	// Implement logout logic
	cookie, err := r.Cookie("refresh_token")
	if err != nil {
		http.Error(w, "no refresh token", http.StatusBadRequest)
		return
	}

	refreshToken := cookie.Value
	claims := &model.Claims{}
	jwt.ParseWithClaims(refreshToken, claims, func(t *jwt.Token) (interface{}, error) {
		return h.JwtSecretKeyRefresh, nil
	})

	// Remove from store
	delete(h.RefreshStore, claims.Email)

	// Clear cookie
	http.SetCookie(w, &http.Cookie{
		Name:     "refresh_token",
		Value:    "",
		Expires:  time.Now().Add(-1 * time.Hour),
		HttpOnly: true,
		Path:     "/",
	})

	w.Write([]byte("logged out"))

}
