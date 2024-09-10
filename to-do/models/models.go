package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model

	UserName string `gorm:"size:24"`
	Password string `gorm:"size:250"`
	Email    string `gorm:"unique;not null"`

	Lists []List
}

type List struct {
	gorm.Model

	Title     string `gorm:"size:24"`
	Completed bool
	UserID    uint `gorm:"not null;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	User User
}
