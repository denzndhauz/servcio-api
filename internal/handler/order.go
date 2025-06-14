package handler

import (
	"net/http"

	"servcio-api/internal/models"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// CreateOrder handles POST /orders
func CreateOrder(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var input struct {
			Items []struct {
				ProductID uint `json:"product_id" binding:"required"`
				Quantity  int  `json:"quantity" binding:"required,min=1"`
			} `json:"items" binding:"required,dive"`
		}

		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		var orderItems []models.OrderItem
		var totalAmount float64

		tx := db.Begin() // Start transaction

		for _, item := range input.Items {
			var product models.Product
			if err := tx.First(&product, item.ProductID).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusNotFound, gin.H{"error": "Product not found"})
				return
			}

			if product.Stock < item.Quantity {
				tx.Rollback()
				c.JSON(http.StatusBadRequest, gin.H{"error": "Not enough stock for product: " + product.Name})
				return
			}

			// Deduct stock
			product.Stock -= item.Quantity
			if err := tx.Save(&product).Error; err != nil {
				tx.Rollback()
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update product stock"})
				return
			}

			orderItem := models.OrderItem{
				ProductID: product.ID,
				Quantity:  item.Quantity,
				Price:     product.Price,
			}
			orderItems = append(orderItems, orderItem)
			totalAmount += product.Price * float64(item.Quantity)
		}

		// Get user ID from context (you'll need JWT middleware to set it)
		userID, exists := c.Get("user_id")
		if !exists {
			tx.Rollback()
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Unauthorized"})
			return
		}

		// Create the order
		order := models.Order{
			UserID:     userID.(uint),
			Total:      totalAmount,
			OrderItems: orderItems,
		}

		if err := tx.Create(&order).Error; err != nil {
			tx.Rollback()
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create order"})
			return
		}

		tx.Commit()

		c.JSON(http.StatusCreated, gin.H{
			"message": "Order created successfully",
			"order":   order,
		})
	}
}
