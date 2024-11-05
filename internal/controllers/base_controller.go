package controllers

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	baseController := NewBaseController()
	r.GET("/ping", baseController.Ping)
}

type BaseController struct {
}

func NewBaseController() *BaseController {
	return &BaseController{}
}

func (b *BaseController) Ping(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
