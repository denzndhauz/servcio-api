package handler

import (
	"net/http"
	"servcio-api/internal/models"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateService(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Define the structure for the input JSON
		var input struct {
			Name        string  `json:"name" binding:"required"`
			Description *string `json:"description" binding:"required,min=1"`
			Duration    int     `json:"duration" binding:"required,min=1"` // in minutes
			Price       float64 `json:"price" binding:"required,min=0"`
			CategoryID  string  `json:"categoryId" binding:"required,uuid"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		tx := db.Begin() // Start transaction

		var category models.ServiceCategory
		if err := tx.First(&category, input.CategoryID).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusNotFound, gin.H{"error": "Category not found"})
			return
		}

		// Check if the service already exists
		var existingService models.Service
		if err := tx.Where("LOWER(name) = ? AND categoryId = ?", strings.ToLower(input.Name), input.CategoryID).First(&existingService).Error; err == nil {
			tx.Rollback()
			c.JSON(http.StatusConflict, gin.H{"error": "Service already exists"})
			return
		}

		service := models.Service{
			Name:        input.Name,
			Description: input.Description,
			Duration:    input.Duration,
			CategoryID:  category.ID,
			Price:       input.Price,
			Providers:   []models.ServiceProvider{}, // Initialize with empty slice
			Bookings:    []models.Booking{},         // Initialize with empty slice
		}

		if err := tx.Create(&service).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service: " + err.Error()})
			return
		}
	}
}

func ListServices(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var services []models.Service
		if err := db.Find(&services).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve services"})
			return
		}

		c.JSON(http.StatusOK, services)
	}
}
