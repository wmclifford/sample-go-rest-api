package tests

import (
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
)

func TestCreateProduct(t *testing.T) {
	db, mock, err := SetupMockDB()
	assert.NoError(t, err, "Failed to setup mock database")

	repo := repositories.NewProductRepository(db)

	product := &models.Product{
		Name:        "Sample Product",
		Description: "Sample Product Description",
		Price:       100.0,
	}

	// Expectation for creating a product
	mock.ExpectBegin()
	mock.ExpectQuery(`INSERT INTO "products" \("name","description","price","created_at","updated_at"\) VALUES \(\$1,\$2,\$3,\$4,\$5\) RETURNING "id"`).
		WithArgs(product.Name, product.Description, product.Price, sqlmock.AnyArg(), sqlmock.AnyArg()).
		WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))
	mock.ExpectCommit()

	err = repo.CreateProduct(product)
	assert.NoError(t, err, "Failed to create product")
	assert.NoError(t, mock.ExpectationsWereMet(), "Mock expectations were not met")
}

func TestGetAllProducts(t *testing.T) {
	db, mock, err := SetupMockDB()
	assert.NoError(t, err, "Failed to setup mock database")

	repo := repositories.NewProductRepository(db)

	// Sample data
	createdAt := time.Now()
	updatedAt := time.Now()
	products := []models.Product{
		{ID: 1, Name: "Product 1", Description: "Product 1 Description", Price: 10.0, CreatedAt: createdAt, UpdatedAt: updatedAt},
		{ID: 2, Name: "Product 2", Description: "Product 2 Description", Price: 20.0, CreatedAt: createdAt, UpdatedAt: updatedAt},
	}

	// Expectation for retrieving all products
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "created_at", "updated_at"}).
		AddRow(products[0].ID, products[0].Name, products[0].Description, products[0].Price, createdAt, updatedAt).
		AddRow(products[1].ID, products[1].Name, products[1].Description, products[1].Price, createdAt, updatedAt)
	mock.ExpectQuery(`SELECT \* FROM "products"`).WillReturnRows(rows)

	result, err := repo.GetAllProducts()
	assert.NoError(t, err, "Failed to retrieve all products")
	assert.Equal(t, products, result, "Retrieved products do not match")
	assert.NoError(t, mock.ExpectationsWereMet(), "Mock expectations were not met")
}

func TestGetProductByID(t *testing.T) {
	db, mock, err := SetupMockDB()
	assert.NoError(t, err, "Failed to setup mock database")

	repo := repositories.NewProductRepository(db)

	// Sample data
	product := &models.Product{
		ID:          1,
		Name:        "Product",
		Description: "Product Description",
		Price:       100.0,
	}

	// Expectation for retrieving a product by ID
	rows := sqlmock.NewRows([]string{"id", "name", "description", "price", "created_at", "updated_at"}).
		AddRow(product.ID, product.Name, product.Description, product.Price, time.Now(), time.Now())
	mock.ExpectQuery(`SELECT \* FROM "products" WHERE "products"\."id" = \$1 ORDER BY "products"\."id" LIMIT \$2`).
		WithArgs(product.ID, 1).
		WillReturnRows(rows)

	result, err := repo.GetProductById(product.ID)
	assert.NoError(t, err, "Failed to retrieve a product by ID")
	assert.Equal(t, product.ID, result.ID, "Retrieved product ID does not match")
	assert.Equal(t, product.Name, result.Name, "Retrieved product name does not match")
	assert.Equal(t, product.Description, result.Description, "Retrieved product description does not match")
	assert.Equal(t, product.Price, result.Price, "Retrieved product price does not match")
	assert.NoError(t, mock.ExpectationsWereMet(), "Mock expectations were not met")
}
