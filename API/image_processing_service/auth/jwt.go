package auth

import (
	"IPS/auth/model"
	"time"

	"github.com/golang-jwt/jwt/v5"
	_ "github.com/joho/godotenv"
)

func GenerateToken(email string, secret_access []byte, secret_refresh []byte, refreshStore map[string]string) (string, string, error) {
	// Access token (short lived)
	accessClaims := model.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(15 * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, accessClaims).SignedString(secret_access)
	if err != nil {
		return "", "", err
	}

	// Refresh token (long lived)
	refreshClaims := model.Claims{
		Email: email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims).SignedString(secret_refresh)
	if err != nil {
		return "", "", err
	}

	// Store refresh token
	refreshStore[email] = refreshToken

	return accessToken, refreshToken, nil
}

func ValidateJWT(tokenString string, jwtkey_access []byte) (*model.Claims, error) {
	claims := &model.Claims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtkey_access, nil
	})
	if err != nil {
		return nil, err
	}
	if !token.Valid {
		return nil, jwt.ErrSignatureInvalid
	}
	return claims, nil
}
