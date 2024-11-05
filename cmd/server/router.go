package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-api/internal/controllers"
	"go-rest-api/internal/services"
)

func setupRouter(userService services.UserService, productService services.ProductService) *gin.Engine {
	router := gin.Default()

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	productController := controllers.NewProductController(productService)
	baseController := controllers.NewBaseController()

	// Base endpoints
	router.GET("/ping", baseController.Ping)

	// User endpoints
	router.POST("/users/register", userController.RegisterUser)

	// Product endpoints
	router.POST("/products", productController.RegisterProduct)
	router.GET("/products", productController.GetAllProducts)
	router.GET("/products/:id", productController.GetProductById)

	return router
}
