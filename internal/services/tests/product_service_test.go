package tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
	"go-rest-api/internal/services"
)

func TestProductService(t *testing.T) {
	db := SetupTestDB(&models.Product{})
	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)

	t.Run("Register a new product", func(t *testing.T) {
		log.Println("Register a new product")
		product := &models.Product{Name: "Sample Product", Description: "Sample Description", Price: 10.0}
		log.Printf("Attempting to register product: %v\n", product)
		err := productService.Register(product)
		assert.Nil(t, err)

		// Ensure the product ID is correctly assigned.
		assert.NotEqual(t, 0, product.ID)
		t.Logf("Product ID after creation: %d", product.ID)

		savedProduct, err := productService.FindById(product.ID)
		assert.Nil(t, err)
		assert.Equal(t, "Sample Product", savedProduct.Name)
		assert.Equal(t, "Sample Description", savedProduct.Description)
		assert.Equal(t, 10.0, savedProduct.Price)
	})
}
