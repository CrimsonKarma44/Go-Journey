package database

import (
	"gorm.io/gorm"
	"time"
)

const (
	Unsuccessful = "unsuccessful"
	Pending      = "pending"
	Successful   = "successful"
)

type Order struct {
	ID uint `gorm:"primarykey"`

	Amount float64
	Status string `gorm:"type:text;check:status IN ('pending', 'unsuccessful', 'successful');default:'pending'"`

	//Status string `gorm:"type:ENUM('pending', 'unsuccessful', 'successful');default:'pending'"`

	CreatedAt time.Time
	UpdatedAt time.Time
}

func (o *Order) CreateOrder(db *gorm.DB) error {
	return db.Create(&o).Error
}

type OrderList struct {
	ID int `gorm:"primary_key"`

	OrderId uint
	Order   Order

	ProductID uint
	Product   Product

	Quantity  int
	ItemPrice float64

	CreatedAt time.Time
}

func (o *OrderList) Create(db *gorm.DB) error {
	return db.Create(&o).Error
}

type Orderlists []OrderList

func (o *Orderlists) GenerateOrder(db *gorm.DB, cartList CartList, order Order) error {
	for _, item := range cartList {
		temp := OrderList{
			OrderId:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			ItemPrice: item.Product.Price,
		}
		err := temp.Create(db)
		if err != nil {
			return err
		}
		*o = append(*o, temp)
	}
	return db.Error
}
