package database

import (
	"fmt"
	"gorm.io/gorm"
)

// Cart is the model representation of the cart of the ecommerce api
type Cart struct {
	ID uint `gorm:"primary_key"`

	UserID uint
	User   User `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	ProductID uint
	Product   Product `gorm:"constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`

	Quantity int
	//Price    float64
	Discount float64
}

// CRUD operations for the cart

func (cart *Cart) AddToCart(db *gorm.DB) error {
	return db.Create(&cart).Error
}

func (cart *Cart) DeleteCart(db *gorm.DB) error {
	return db.Unscoped().Delete(&cart).Error
}

type CartList []Cart

func (c *CartList) Get(db *gorm.DB, userID uint) error {
	return db.Preload("Product").Preload("Product.User").Preload("User").Where("user_id = ?", userID).Find(&c).Error
}

func (c *CartList) TotalPrice() float64 {
	totalPrice := 0.0
	for _, cart := range *c {
		fmt.Println(cart.Product.Price)
		fmt.Println(cart.Discount)
		fmt.Println(cart.Quantity)
		totalPrice += (cart.Product.Price - (cart.Product.Price * cart.Discount)) * float64(cart.Quantity)
	}
	return totalPrice
}
