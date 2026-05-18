package handler

import (
	"ecommerce/database"
	"encoding/json"
	"fmt"
	"gorm.io/gorm"
	"net/http"
	"strconv"
)

type ProductHandler struct {
	DB *gorm.DB
}

func (handler *ProductHandler) CreateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		fmt.Println(r.FormValue("name"))
		productItem := database.Product{
			Name:        r.FormValue("name"),
			Description: r.FormValue("description"),
			Price: func() float64 {
				value, err := strconv.ParseFloat(r.FormValue("price"), 64)
				fmt.Println(value, r.FormValue("price"))
				if err != nil {
					return 0
				}
				return value
			}(),
			Quantity: func() int {
				value, err := strconv.Atoi(r.FormValue("quantity"))
				if err != nil {
					return 0
				}
				return value
			}(),
			ImageURL: r.FormValue("image_url"),
			UserID: func() uint {
				value, err := strconv.Atoi(r.FormValue("user_id"))
				if err != nil {
					return 0
				}
				return uint(value)
			}(),
		}
		err := productItem.Create(handler.DB)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.WriteHeader(http.StatusOK)
		w.Write([]byte("Product created successfully"))
		return
	}
}

func (handler *ProductHandler) ViewHandler(w http.ResponseWriter, r *http.Request) {
	var productItem database.Product
	if r.Method == http.MethodGet {
		productID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		productItem.Get(handler.DB, uint(productID))
		w.WriteHeader(http.StatusOK)
		jsonText, err := json.Marshal(productItem)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Write(jsonText)
		return
	}
}

func (handler *ProductHandler) DeleteHandler(w http.ResponseWriter, r *http.Request) {
	var productItem database.Product
	if r.Method == http.MethodDelete {
		productID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		err = productItem.Get(handler.DB, uint(productID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		userId, err := strconv.Atoi(r.FormValue("user_id"))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if uint(userId) != productItem.UserID {
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		err = productItem.Delete(handler.DB, uint(productID))
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
	}
}
