package handler

import (
	"net/http"
	"servcio-api/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateCategory(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Name        string  `json:"name" binding:"required"`
			Description *string `json:"description"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tx := db.Begin() // Start transaction

		var existingCategory models.ServiceCategory
		if err := tx.Where("LOWER(name) = ?", strings.ToLower(input.Name)).First(&existingCategory).Error; err == nil {
			tx.Rollback()
			c.JSON(http.StatusConflict, gin.H{"error": "Category already exists"})
			return
		}

		category := models.ServiceCategory{
			Name:        input.Name,
			Description: input.Description,
		}

		if err := tx.Create(&category).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create category"})
			return
		}

		tx.Commit()
		c.JSON(http.StatusCreated, category)
	}
}

func ListCategories(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var categories []models.ServiceCategory
		if err := db.Find(&categories).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve categories"})
			return
		}

		c.JSON(http.StatusOK, categories)
	}
}
