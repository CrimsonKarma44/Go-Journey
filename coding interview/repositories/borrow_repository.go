package repositories

import (
	"coding_interview/models"
	"gorm.io/gorm"
	_ "time"
)

type BorrowRepository struct {
	DB *gorm.DB
}

func (r *BorrowRepository) BorrowBook(borrow *models.Borrow) error {
	return r.DB.Create(borrow).Error
}

func (r *BorrowRepository) UpdateBookStatus(bookID uint, isBorrowed bool) error {
	return r.DB.Model(&models.Book{}).Where("id = ?", bookID).Update("is_borrowed", isBorrowed).Error
}
