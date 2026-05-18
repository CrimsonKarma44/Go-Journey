package model

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type Creds struct {
	gorm.Model
	Email    string
	Password string
}

func (u *Creds) validate() error {
	if u.Email == "" {
		return nil
	}
	if u.Password == "" {
		return nil
	}
	return nil
}

func (u *Creds) Encrypt() error {
	if err := u.validate(); err != nil {
		return err
	}
	byctes, err := bcrypt.GenerateFromPassword([]byte(u.Password), 14)
	if err != nil {
		return err
	}
	u.Password = string(byctes)
	return nil
}

func (u *Creds) CheckPassword(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err == nil
}
