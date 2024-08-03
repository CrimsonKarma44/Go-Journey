package main

import (
	"fmt"
	"to-do/database"
	"to-do/utils"
)

func main() {
	db, err := database.GormPostgresqlInit()
	if err != nil {
		panic(err)
	}
	//err = utils.InsertingUsers(db, &models.User{Email: "vincentprincewill44@yahoo.com", UserName: "Karma", Password: "password"})
	users := utils.RetrievingAllUser(db)
	//fmt.Println(users)
	for _, user := range users {
		fmt.Println(user.Email)
	}
	//if err != nil {
	//	panic(err)
	//}
	//result := db.Preload("Lists").Find(&dummy)
	//err = utils.DeletingUser(db, "vincentprincewill44@gmail.com")
	//if err != nil {
	//	panic(err)
	//} else {
	//	fmt.Println("Deleted successfully")
	//}
	//fmt.Println(utils.RetrievingAllUser(db))
	//var user_ []models.User
	//for _, user := range users {
	//	fmt.Println(user.UserID)
	//	db.Find(&user_, "id = ?", user.UserID)
	//	for _, u := range user_ {
	//		fmt.Println(u.UserName)
	//	}
}
