package model

import (
	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
	Email string `json:"username"`
	jwt.RegisteredClaims
}
