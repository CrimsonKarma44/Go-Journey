package services

import (
	"coding_interview/models"
	"coding_interview/repositories"
)

type BookService struct {
	BookRepo *repositories.BookRepository
}

func (s *BookService) ListBooks() ([]models.Book, error) {
	return s.BookRepo.ListBooks()
}

func (s *BookService) GetBookByID(id uint) (*models.Book, error) {
	return s.BookRepo.GetBookByID(id)
}

func (s *BookService) FilterBooks(publisher, category string) ([]models.Book, error) {
	return s.BookRepo.FilterBooks(publisher, category)
}
