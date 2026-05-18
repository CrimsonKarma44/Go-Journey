package database

import (
	"ecommerce/utils"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	Username string
	Password string
	Email    string `gorm:"unique"`
	Role     string
}

func (u *User) Create(db *gorm.DB) error {
	var err error
	u.Password, err = utils.HashPassword(u.Password)
	if err != nil {
		return err
	}
	tx := db.Create(&u)
	return tx.Error
}

func (u *User) RetrieveEmail(db *gorm.DB, email string) error {
	return db.Where("email = ?", email).First(&u).Error
}
func (u *User) RetrieveID(db *gorm.DB, id uint) error {
	return db.Where("id = ?", id).First(&u).Error
}

func (u *User) Update(db *gorm.DB) error {
	return db.Save(&u).Error
}
