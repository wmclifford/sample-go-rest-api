package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"go-rest-api/internal/controllers"
	"go-rest-api/internal/models"
	"go-rest-api/internal/services/tests"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Unique setup function
func setupUserTestEnv(t *testing.T) (*gin.Engine, *tests.MockUserService, *controllers.UserController) {
	mockUserService := new(tests.MockUserService)
	controller := controllers.NewUserController(mockUserService)

	router := gin.Default()
	return router, mockUserService, controller
}

// Reuse the common executeRequest function
func executeUserRequest(router *gin.Engine, req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	router.ServeHTTP(rr, req)
	return rr
}

func TestRegisterUser(t *testing.T) {
	router, mockUserService, controller := setupUserTestEnv(t)
	router.POST("/users/register", controller.RegisterUser)

	// Create a user payload
	user := models.User{
		Name:  "John Doe",
		Email: "john.doe@example.com",
	}
	userJSON, _ := json.Marshal(user)

	// Setup expected behavior from the mock service
	mockUserService.On("Register", mock.AnythingOfType("*models.User")).Return(nil)

	req, _ := http.NewRequest("POST", "/users/register", bytes.NewBuffer(userJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := executeUserRequest(router, req)

	assert.Equal(t, http.StatusOK, resp.Code)
	mockUserService.AssertExpectations(t)
}

// Additional test cases can be added in a similar manner
