package main

import (
	"fmt"
	"to-do/database"
	"to-do/utils"
)

func main() {
	//fmt.Println(database.DNS())
	db, err := database.GormInit()
	if err != nil {
		panic(err)
	}
	//result := db.Preload("Lists").Find(&dummy)
	err = utils.DeletingUser(db, "vincentprincewill44@gmail.com")
	if err != nil {
		panic(err)
	} else {
		fmt.Println("Deleted successfully")
	}
	fmt.Println(utils.RetrievingAllUser(db))
	//var user_ []models.User
	//for _, user := range users {
	//	fmt.Println(user.UserID)
	//	db.Find(&user_, "id = ?", user.UserID)
	//	for _, u := range user_ {
	//		fmt.Println(u.UserName)
	//	}
}
