package handler

import (
	"net/http"
	"servcio-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateServiceProvider(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			FirstName   string   `json:"first_name" binding:"required"`
			LastName    string   `json:"last_name" binding:"required"`
			Email       string   `json:"email" binding:"required,email"`
			Phone       *string  `json:"phone"`
			Specialties []string `json:"specialties"` // List of service IDs
			Schedules   []string `json:"schedules"`   // List of schedule IDs
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		// Convert input.Specialties (IDs) to []models.Service
		var specialties []models.Service
		if len(input.Specialties) > 0 {
			if err := db.Where("id IN ?", input.Specialties).Find(&specialties).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch specialties"})
				return
			}
		}

		var schedules []models.Schedule
		if len(input.Schedules) > 0 {
			if err := db.Where("id IN ?", input.Schedules).Find(&schedules).Error; err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch schedules"})
				return
			}
		}

		// Check email and phone uniqueness
		var existingProvider models.ServiceProvider
		if err := db.Where("email = ? OR phone = ?", input.Email, input.Phone).First(&existingProvider).Error; err == nil {
			c.JSON(http.StatusConflict, gin.H{"error": "Service provider with given email or phone already exists"})
			return
		}

		tx := db.Begin() // Start transaction

		provider := models.ServiceProvider{
			FirstName:   input.FirstName,
			LastName:    input.LastName,
			Email:       input.Email,
			Phone:       input.Phone,
			Specialties: specialties,
			Schedules:   schedules,
		}

		if err := db.Create(&provider).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create service provider"})
			return
		}

		tx.Commit()
		c.JSON(http.StatusCreated, provider)
	}
}

func ListServiceProviders(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var providers []models.ServiceProvider
		if err := db.Preload("Specialties").Preload("Schedules").Find(&providers).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve service providers" + err.Error()})
			return
		}
		c.JSON(http.StatusOK, providers)
	}
}
