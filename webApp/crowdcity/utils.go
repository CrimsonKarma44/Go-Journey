package main

import (
	"github.com/dgrijalva/jwt-go"
	"os"
	"time"
)

// Claims is a custom JWT claims structure
type Claims struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.StandardClaims
}

// GenerateJWT generates a new JWT token for a given username.
func GenerateJWT(username string, email string) (string, error) {
	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		Username: username,
		Email:    email,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ValidateJWT validates a given JWT token and returns the claims if valid.
func ValidateJWT(tokenString string) (*Claims, error) {
	claims := &Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		return nil, err
	}

	if !token.Valid {
		return nil, jwt.NewValidationError("invalid token", jwt.ValidationErrorExpired)
	}

	return claims, nil
}
