package models

import (
	"gorm.io/gorm"
	"time"
)

type Borrow struct {
	gorm.Model
	UserID   uint      `gorm:"not null"`
	BookID   uint      `gorm:"not null"`
	DueDate  time.Time `gorm:"not null"`
	Returned bool      `gorm:"default:false"`
}
