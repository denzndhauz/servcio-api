package handler

import (
	"net/http"
	"servcio-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateProduct(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input models.Product
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if err := db.Create(&input).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create product"})
			return
		}

		c.JSON(http.StatusCreated, input)
	}
}

func ListProducts(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var products []models.Product
		if err := db.Find(&products).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve products"})
			return
		}

		c.JSON(http.StatusOK, products)
	}
}
