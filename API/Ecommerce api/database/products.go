package database

import (
	"gorm.io/gorm"
)

type Product struct {
	gorm.Model
	Name        string
	Description string
	Price       float64
	Quantity    int
	ImageURL    string

	UserID uint
	User   User
}

func (p *Product) Create(db *gorm.DB) error {
	return db.Create(&p).Error
}

func (p *Product) Get(db *gorm.DB, id uint) error {
	return db.Preload("User").Where("id = ?", id).First(&p).Error
}

func (p *Product) Update(db *gorm.DB) error {
	return db.Save(&p).Error
}

func (p *Product) Delete(db *gorm.DB, id uint) error {
	return db.Unscoped().Delete(&p, id).Error
}

type ProductsList []Product

func (pl *ProductsList) Get(db *gorm.DB, userId string) error {
	return db.Where("user_id = ?", userId).Find(pl).Error
}
