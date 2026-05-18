package repositories

import (
	"coding_interview/models"
	"gorm.io/gorm"
)

type BookRepository struct {
	DB *gorm.DB
}

func (r *BookRepository) CreateBook(book *models.Book) error {
	return r.DB.Create(book).Error
}

func (r *BookRepository) GetBookByID(id uint) (*models.Book, error) {
	var book models.Book
	err := r.DB.First(&book, id).Error
	return &book, err
}

func (r *BookRepository) ListBooks() ([]models.Book, error) {
	var books []models.Book
	err := r.DB.Find(&books).Error
	return books, err
}

func (r *BookRepository) FilterBooks(publisher, category string) ([]models.Book, error) {
	var books []models.Book
	query := r.DB
	if publisher != "" {
		query = query.Where("publisher = ?", publisher)
	}
	if category != "" {
		query = query.Where("category = ?", category)
	}
	err := query.Find(&books).Error
	return books, err
}
