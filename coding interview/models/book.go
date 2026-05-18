package models

import "gorm.io/gorm"

type Book struct {
	gorm.Model
	Title      string `gorm:"not null"`
	Author     string `gorm:"not null"`
	Publisher  string `gorm:"not null"`
	Category   string `gorm:"not null"`
	IsBorrowed bool   `gorm:"default:false"`
}

//package models
//
//import "gorm.io/gorm"
//
//type Book struct {
//	gorm.Model
//
//	Title        string
//	Author       string
//	Availability bool
//
//	Publisher string
//	Category  string
//}
//
//func (book *Book) RetrieveID(db gorm.DB, id uint) error {
//	return db.Where("id = ?", id).First(book).Error
//}
//
//type Books []Book
//
//func (books Books) GetAvailBooks(db *gorm.DB) error {
//	return db.Where("availability = ?", true).Find(&books).Error
//}
//
//func (books Books) FilterBooksPub(db *gorm.DB, publisher string) error {
//	return db.Where("publisher = ?", publisher).Find(&books).Error
//}
//func (books Books) FilterBooksCat(db *gorm.DB, Cat string) error {
//	return db.Where("category = ?", Cat).Find(&books).Error
//}
//
//type Borrow struct {
//	gorm.Model
//
//	UserID uint
//	user   User
//}
