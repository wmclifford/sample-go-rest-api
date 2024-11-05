package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"go-rest-api/internal/controllers"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
	"go-rest-api/internal/services/tests"
)

// Unique setup function
func setupProductTestEnv(t *testing.T) (*gin.Engine, *tests.MockProductService, *controllers.ProductController) {
	mockProductService := new(tests.MockProductService)
	controller := controllers.NewProductController(mockProductService)

	router := gin.Default()
	return router, mockProductService, controller
}

// Reuse the common executeRequest function
func executeProductRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func TestRegisterProduct(t *testing.T) {
	router, mockProductService, controller := setupProductTestEnv(t)
	router.POST("/products", controller.RegisterProduct)

	// Create a product payload
	product := models.Product{
		Name:  "Sample Product",
		Price: 100,
	}
	productJSON, _ := json.Marshal(product)

	// Setup expected behavior from the mock service
	mockProductService.On("Register", mock.AnythingOfType("*models.Product")).Return(nil)

	req, _ := http.NewRequest("POST", "/products", bytes.NewBuffer(productJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := executeProductRequest(router, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Status code should be 200")
	mockProductService.AssertExpectations(t)
}

func TestGetAllProducts(t *testing.T) {
	router, mockProductService, controller := setupProductTestEnv(t)
	router.GET("/products", controller.GetAllProducts)

	products := []models.Product{
		{Name: "Product 1", Price: 10},
		{Name: "Product 2", Price: 20},
	}

	// Setup expected behavior from the mock service
	mockProductService.On("FindAll").Return(products, nil)

	req, _ := http.NewRequest("GET", "/products", nil)
	resp := executeProductRequest(router, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Status code should be 200")
	var returnedProducts []models.Product
	err := json.Unmarshal(resp.Body.Bytes(), &returnedProducts)
	assert.Nil(t, err, "Error unmarshalling response body")
	assert.Equal(t, products, returnedProducts, "Returned products should match")
	mockProductService.AssertExpectations(t)
}

func TestGetProductById(t *testing.T) {
	router, mockProductService, controller := setupProductTestEnv(t)
	router.GET("/products/:id", controller.GetProductById)

	product := models.Product{ID: 1, Name: "Product 1", Price: 10}

	// Setup expected behavior from the mock service
	mockProductService.On("FindById", uint(1)).Return(&product, nil)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	resp := executeProductRequest(router, req)

	assert.Equal(t, http.StatusOK, resp.Code, "Status code should be 200")
	var returnedProduct models.Product
	err := json.Unmarshal(resp.Body.Bytes(), &returnedProduct)
	assert.Nil(t, err, "Error unmarshalling response body")
	assert.Equal(t, product, returnedProduct, "Returned product should match")
	mockProductService.AssertExpectations(t)
}

func TestGetProductByID_InvalidID(t *testing.T) {
	router, _, controller := setupProductTestEnv(t)
	router.GET("/products/:id", controller.GetProductById)

	req, _ := http.NewRequest("GET", "/products/invalid-id", nil)
	resp := executeProductRequest(router, req)

	assert.Equal(t, http.StatusBadRequest, resp.Code, "Status code should be 400")
	assert.Contains(t, resp.Body.String(), "invalid ID format", "Response body should contain error message")
}

func TestGetProductByID_ProductNotFound(t *testing.T) {
	router, mockProductService, controller := setupProductTestEnv(t)
	router.GET("/products/:id", controller.GetProductById)

	// Setup expected behavior from the mock service
	mockProductService.On("FindById", uint(1)).Return(nil, repositories.ErrProductNotFound)

	req, _ := http.NewRequest("GET", "/products/1", nil)
	resp := executeProductRequest(router, req)

	assert.Equal(t, http.StatusNotFound, resp.Code, "Status code should be 404")
	mockProductService.AssertExpectations(t)
}
