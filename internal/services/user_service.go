package services

import (
	"errors"

	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
)

type UserService interface {
	Register(user *models.User) error
	FindByEmail(email string) (*models.User, error)
}

type userService struct {
	repo repositories.UserRepository
}

func NewUserService(repo repositories.UserRepository) UserService {
	return &userService{repo: repo}
}

func (s *userService) Register(user *models.User) error {
	// may want to hash the password before saving ...

	if existingUser, err := s.repo.GetUserByEmail(user.Email); err == nil && existingUser != nil {
		return errors.New("user already exists")
	}
	return s.repo.CreateUser(user)
}

func (s *userService) FindByEmail(email string) (*models.User, error) {
	return s.repo.GetUserByEmail(email)
}
