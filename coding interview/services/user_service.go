package services

import (
	"coding_interview/models"
	"coding_interview/repositories"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func (s *UserService) EnrollUser(email, firstName, lastName string) (*models.User, error) {
	user := &models.User{
		Email:     email,
		FirstName: firstName,
		LastName:  lastName,
	}
	err := s.UserRepo.CreateUser(user)
	return user, err
}
