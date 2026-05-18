package models

import (
	"gorm.io/gorm"
	"image"
	"verhiser/utils"
)

type User struct {
	gorm.Model
	Username string `gorm:"unique" json:"username"`
	Email    string `json:"email" gorm:"unique"`
	Password string `json:"password"`
	Class    string `json:"class"`
}

func (u *User) Create(db *gorm.DB) (*gorm.DB, error) {
	u.Password, _ = utils.HashPassword(u.Password)

	err := db.Create(&u)
	if err != nil {
		//panic(err)
		return nil, err.Error
	}
	return db, nil
}

func (u *User) Get(db *gorm.DB) error {
	err := db.Where(&User{Username: u.Username}).First(&u).Error
	if err != nil {
		return err
	}
	return nil
}
