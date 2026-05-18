package models

import "gorm.io/gorm"

type User struct {
	gorm.Model

	Email     string `gorm:"unique;not null"`
	FirstName string `gorm:"not null"`
	LastName  string `gorm:"not null"`
}

//type User struct {
//	gorm.Model
//
//	FirstName string
//	LastName  string
//	Email     string
//	Password  string
//}
//
//func (u *User) Create(db *gorm.DB) error {
//	var err error
//	u.Password, err = utils.HashPassword(u.Password)
//	if err != nil {
//		return err
//	}
//	tx := db.Create(&u)
//	return tx.Error
//}
