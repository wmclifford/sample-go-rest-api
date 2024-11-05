package repositories

import (
	"errors"
	"log"

	"go-rest-api/internal/models"
	"gorm.io/gorm"
)

// ProductRepository defines the methods that any data storage provider needs
// to implement to get and store product data.
type ProductRepository interface {
	CreateProduct(product *models.Product) error
	GetAllProducts() ([]models.Product, error)
	GetProductById(id uint) (*models.Product, error)
}

type productRepository struct {
	db *gorm.DB
}

// NewProductRepository creates a new instance of ProductRepository.
func NewProductRepository(db *gorm.DB) ProductRepository {
	return &productRepository{db: db}
}

// CreateProduct inserts a new Product into the database.
func (r *productRepository) CreateProduct(product *models.Product) error {
	log.Printf("creating product: %v", product)
	result := r.db.Create(product)
	if result.Error != nil {
		log.Printf("error creating product: %v", result.Error)
	} else {
		log.Printf("created product: %v", product)
	}
	return result.Error
}

// GetAllProducts retrieves all Products from the database.
func (r *productRepository) GetAllProducts() ([]models.Product, error) {
	var products []models.Product
	result := r.db.Find(&products)
	if result.Error != nil {
		log.Printf("error finding all products: %v", result.Error)
	}
	return products, result.Error
}

var ErrProductNotFound = errors.New("product not found")

// GetProductById retrieves a Product by its ID from the database.
func (r *productRepository) GetProductById(id uint) (*models.Product, error) {
	var product models.Product
	result := r.db.First(&product, id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			log.Printf("product not found by ID: %d", id)
			return nil, ErrProductNotFound
		}
		log.Printf("error finding product by ID: %d, error: %v", id, result.Error)
	} else {
		log.Printf("found product: %v", product)
	}
	return &product, result.Error
}
