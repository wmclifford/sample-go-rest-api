package services

import (
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
)

type ProductService interface {
	Register(product *models.Product) error
	FindAll() ([]models.Product, error)
	FindById(id uint) (*models.Product, error)
}

type productService struct {
	repo repositories.ProductRepository
}

func NewProductService(repo repositories.ProductRepository) ProductService {
	return &productService{repo: repo}
}

func (s *productService) Register(product *models.Product) error {
	return s.repo.CreateProduct(product)
}

func (s *productService) FindAll() ([]models.Product, error) {
	return s.repo.GetAllProducts()
}

func (s *productService) FindById(id uint) (*models.Product, error) {
	return s.repo.GetProductById(id)
}
