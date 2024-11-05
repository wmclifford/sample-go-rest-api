package tests

import (
	"github.com/stretchr/testify/mock"
	"go-rest-api/internal/models"
)

type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Register(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *MockUserService) FindByEmail(email string) (*models.User, error) {
	args := m.Called(email)
	return args.Get(0).(*models.User), args.Error(1)
}

// Implement other methods if needed
