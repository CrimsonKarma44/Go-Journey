package handler

import (
	"ecommerce/database"
	"encoding/json"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type CartHandler struct {
	DB *gorm.DB
}

func (handler *CartHandler) AddToCart(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		cart := database.Cart{
			UserID: func() uint {
				v, err := strconv.Atoi(r.FormValue("user_id"))
				if err != nil {
					return 0
				}
				return uint(v)
			}(),
			ProductID: func() uint {
				v, err := strconv.Atoi(r.FormValue("product_id"))
				if err != nil {
					return 0
				}
				return uint(v)
			}(),
			Quantity: func() int {
				v, err := strconv.Atoi(r.FormValue("quantity"))
				if err != nil {
					return 0
				}
				return v
			}(),
			Discount: func() float64 {
				v, err := strconv.ParseFloat(r.FormValue("discount"), 64)
				if err != nil {
					return 0
				}
				return v
			}(),
		}
		err := cart.AddToCart(handler.DB)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte(`{"message":"Product successfully added to Cart"}`))
		jsoned, _ := json.Marshal(cart)
		w.Write(jsoned)
		return
	}
}

func (handler *CartHandler) ViewCart(w http.ResponseWriter, r *http.Request) {
	var cart database.CartList
	if r.Method == http.MethodGet {
		userId, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		err = cart.Get(handler.DB, uint(userId))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		jsoned, _ := json.Marshal(cart)
		w.Write(jsoned)
		return
	}
}

func (handler *CartHandler) RemoveFromCart(w http.ResponseWriter, r *http.Request) {}
