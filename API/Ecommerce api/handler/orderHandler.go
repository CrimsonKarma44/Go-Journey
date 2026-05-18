package handler

//type OrderHandler struct {
//	DB *gorm.DB
//}
//
//func (o *OrderHandler) Order(w http.ResponseWriter, r *http.Request) {
//	if r.Method == http.MethodPost {
//		userCart := database.CartList{}
//		err := userCart.Get(o.DB,
//			func() uint {
//				value, err := strconv.Atoi(r.FormValue("id"))
//				if err != nil {
//					w.WriteHeader(http.StatusBadRequest)
//				}
//				return uint(value)
//			}())
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//		order := database.Order{
//			Amount: userCart.TotalPrice(),
//		}
//		_, err = urls.GenerateOrder(o.DB, userCart, order)
//		if err != nil {
//			w.WriteHeader(http.StatusBadRequest)
//			return
//		}
//jsonOrder, err := json.Marshal(orderList)
//if err != nil {
//	w.WriteHeader(http.StatusBadRequest)
//	return
//}
//w.Header().Set("Content-Type", "application/json")
//		w.WriteHeader(http.StatusOK)
//		w.Write(jsonOrder)
//		return
//	}
//}
