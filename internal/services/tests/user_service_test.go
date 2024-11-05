package tests

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
	"go-rest-api/internal/services"
)

func TestUserService(t *testing.T) {
	db := SetupTestDB(&models.User{})
	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	t.Run("Register a new user", func(t *testing.T) {
		log.Println("Register a new user")
		user := &models.User{Name: "John Doe", Email: "johndoe@example.com", Password: "securepassword"}
		log.Printf("Attempting to register user: %v\n", user)
		err := userService.Register(user)
		assert.Nil(t, err)

		// Ensure the user ID is correctly assigned.
		assert.NotEqual(t, 0, user.ID)
		t.Logf("User ID after creation: %d\n", user.ID)

		log.Println("Verifying user was saved")
		savedUser, err := userService.FindByEmail("johndoe@example.com")
		assert.Nil(t, err)
		assert.Equal(t, "John Doe", savedUser.Name)
		assert.Equal(t, "johndoe@example.com", savedUser.Email)
	})

	t.Run("Fail to register an existing user", func(t *testing.T) {
		user := &models.User{Name: "John Doe", Email: "johndoe@example.com", Password: "securepassword"}
		err := userService.Register(user)
		assert.NotNil(t, err)
		assert.Equal(t, "user already exists", err.Error())
	})
}
