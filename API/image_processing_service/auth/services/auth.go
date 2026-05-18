package services

import (
	"IPS/auth/model"
	"encoding/json"
	"fmt"

	"gorm.io/gorm"
)

type AuthService struct {
	DB *gorm.DB
}

func (a AuthService) New() AuthService {
	return AuthService{DB: a.DB}
}

func (a *AuthService) Register(user *model.Creds) ([]byte, error) {
	if r := a.DB.Where("email = ?", user.Email).First(&user); r.Error == nil {
		return nil, fmt.Errorf("User already exists")
	}

	if err := user.Encrypt(); err != nil {
		return nil, err
	}

	result := a.DB.Create(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	response, _ := json.Marshal(map[string]interface{}{
		"id":        user.ID,
		"email":     user.Email,
		"createdAt": user.CreatedAt,
	})

	return response, nil
}

func (a *AuthService) Login(user model.Creds) ([]byte, error) {
	var storedUser model.Creds

	result := a.DB.Where("email = ?", user.Email).First(&storedUser)
	if result.Error != nil {
		return nil, result.Error
	}

	if !storedUser.CheckPassword(user.Password) {
		return nil, fmt.Errorf("invalid password")
	}

	response, _ := json.Marshal(map[string]interface{}{
		"id":        storedUser.ID,
		"email":     storedUser.Email,
		"createdAt": storedUser.CreatedAt,
	})

	return response, nil
}

func (a *AuthService) Logout(userID uint) error {
	// Implement logout logic if needed (e.g., token invalidation)
	return nil
}
