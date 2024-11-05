package repositories

import (
	"log"

	"go-rest-api/internal/models"
	"gorm.io/gorm"
)

// UserRepository defines the interface for user data operations.
type UserRepository interface {
	// CreateUser creates a new user record in the database.
	CreateUser(user *models.User) error

	// GetUserByEmail retrieves a user record by email from the database.
	GetUserByEmail(email string) (*models.User, error)
}

// userRepository implements the UserRepository interface.
type userRepository struct {
	db *gorm.DB
}

// NewUserRepository creates a new instance of the userRepository.
func NewUserRepository(db *gorm.DB) UserRepository {
	return &userRepository{db: db}
}

// CreateUser creates a new user record in the database.
func (r *userRepository) CreateUser(user *models.User) error {
	log.Printf("creating user: %v", user)
	result := r.db.Create(user)
	if result.Error != nil {
		log.Printf("error creating user: %v", result.Error)
	} else {
		log.Printf("created user: %v", user)
	}
	return result.Error
}

// GetUserByEmail retrieves a user record by email from the database.
func (r *userRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	result := r.db.Where("email = ?", email).First(&user)
	if result.Error != nil {
		log.Printf("error finding user by email: %s, error: %v", email, result.Error)
	} else {
		log.Printf("found user: %v", user)
	}
	return &user, result.Error
}
