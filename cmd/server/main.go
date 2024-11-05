package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"time"

	"go-rest-api/config"
	"go-rest-api/internal/models"
	"go-rest-api/internal/repositories"
	"go-rest-api/internal/services"
	"go-rest-api/pkg/database"
	"gorm.io/gorm"
)

const maxRetries = 10
const retryInterval = 5 * time.Second

func main() {
	// Get the current working directory
	wd, err := os.Getwd()
	if err != nil {
		fmt.Printf("Could not get current working directory: %v\n", err)
		return
	}

	// Build the path to the config file
	configFile := filepath.Join(wd, "config", "config.yaml")

	// Load configuration
	if err = config.Load(configFile); err != nil {
		fmt.Printf("Failed to load configuration: %v\n", err)
		return
	}

	// Retry mechanism for database connection
	var db *gorm.DB
	for i := 0; i < maxRetries; i++ {
		db, err = database.InitDB(database.GetDSN())
		if err == nil {
			break
		}
		fmt.Printf("Could not connect to database: %v\n", err)
		fmt.Printf("Retrying in %s...\n", retryInterval)
		time.Sleep(retryInterval)
	}

	if err != nil {
		log.Fatalf("Could not connect to database after %d attempts: %v", maxRetries, err)
	}

	// Ensure both User and Product models are migrated
	err = db.AutoMigrate(&models.User{}, &models.Product{})
	if err != nil {
		log.Fatalf("Could not migrate database: %v", err)
	}

	userRepo := repositories.NewUserRepository(db)
	userService := services.NewUserService(userRepo)

	productRepo := repositories.NewProductRepository(db)
	productService := services.NewProductService(productRepo)

	// Setup Gin router
	r := setupRouter(userService, productService)

	// Run the server
	err = r.Run(":8080")
	if err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
