package utils

import (
	"fmt"
	"gorm.io/gorm"
	"to-do/models"
)

func RetrievingAllUser(db *gorm.DB) []models.User {
	var users []models.User
	db.Find(&users)
	return users
}

func RetrievingAllList(db *gorm.DB) []models.List {
	var lists []models.List
	db.Find(&lists)
	return lists
}

func InsertingList(db *gorm.DB, lists *models.List) error {
	results := db.Create(lists)
	if results.Error != nil {
		return results.Error
	}
	return nil
}
func InsertingUsers(db *gorm.DB, users *models.User) error {
	result := db.Create(users)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func DeletingUser(db *gorm.DB, email string) error {
	var users models.User
	db.First(&users, "email = ?", email)
	fmt.Println(users)
	result := db.Delete(&users)
	if result.Error != nil {
		fmt.Println("err on delete")
		return result.Error
	}
	return nil
}
func DeletingList(db *gorm.DB, list *models.User) error {
	result := db.Delete(list)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
