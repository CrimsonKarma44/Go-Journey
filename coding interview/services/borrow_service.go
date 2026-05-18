package services

import (
	"coding_interview/models"
	"coding_interview/repositories"
	"time"
)

type BorrowService struct {
	BorrowRepo *repositories.BorrowRepository
	BookRepo   *repositories.BookRepository
}

func (s *BorrowService) BorrowBook(userID, bookID uint, days int) error {
	dueDate := time.Now().AddDate(0, 0, days)
	borrow := &models.Borrow{
		UserID:  userID,
		BookID:  bookID,
		DueDate: dueDate,
	}
	err := s.BorrowRepo.BorrowBook(borrow)
	if err != nil {
		return err
	}
	return s.BorrowRepo.UpdateBookStatus(bookID, true)
}
