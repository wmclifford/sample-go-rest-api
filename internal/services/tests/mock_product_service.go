package tests

import (
	"github.com/stretchr/testify/mock"
	"go-rest-api/internal/models"
)

type MockProductService struct {
	mock.Mock
}

func (m *MockProductService) Register(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *MockProductService) FindAll() ([]models.Product, error) {
	args := m.Called()
	return args.Get(0).([]models.Product), args.Error(1)
}

func (m *MockProductService) FindById(id uint) (*models.Product, error) {
	args := m.Called(id)
	product := args.Get(0)
	if product == nil {
		return nil, args.Error(1)
	}
	return product.(*models.Product), args.Error(1)
}

// Implement other ProductService methods if needed
