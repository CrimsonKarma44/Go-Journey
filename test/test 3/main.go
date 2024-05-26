package main

import (
	"encoding/json"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"time"
)

var jwtKey = []byte("your_secret_key")

// Credentials struct to represent the user credentials
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Claims struct to represent the JWT claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func main() {
	http.HandleFunc("/login", login)
	http.HandleFunc("/dd", retrieve)
	http.HandleFunc("/protected", protected)

	fmt.Println("Server is running on http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}

func retrieve(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return JWT token as response
	w.Write([]byte(tokenString))
}

func login(w http.ResponseWriter, r *http.Request) {
	var creds Credentials

	// Parse JSON request body
	if err := json.NewDecoder(r.Body).Decode(&creds); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Authenticate the user (in this example, just compare hardcoded values)
	if creds.Username != "user" || creds.Password != "password" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Generate JWT token
	expirationTime := time.Now().Add(100 * time.Minute)
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Return JWT token as response
	w.Write([]byte(tokenString))
}

func protected(w http.ResponseWriter, r *http.Request) {
	// Extract JWT token from request headers
	tokenString := r.Header.Get("Authorization")
	if tokenString == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Parse JWT token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Check if the token is valid
	if !token.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Extract claims from token
	claims, ok := token.Claims.(*Claims)
	if !ok {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	// Do something with the claims (e.g., authenticate user, authorize access)
	// In this example, we just return a success message
	w.Write([]byte(fmt.Sprintf("Welcome %s!", claims.Username)))
}
