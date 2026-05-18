package urls

import (
	"ecommerce/database"
	"gorm.io/gorm"
)

func GenerateOrder(db *gorm.DB, cartList database.CartList, order database.Order) (database.Orderlists, error) {
	list := database.Orderlists{}
	for _, item := range cartList {
		temp := database.OrderList{
			OrderId:   order.ID,
			ProductID: item.ProductID,
			Quantity:  item.Quantity,
			ItemPrice: item.Product.Price,
		}
		err := temp.Create(db)
		if err != nil {
			return nil, err
		}
		list = append(list, temp)
	}
	return list, db.Error
}
