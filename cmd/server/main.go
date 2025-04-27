package main

import (
	"ecomprac/internal/config"
	"ecomprac/internal/database"
	"ecomprac/internal/handler"
	"ecomprac/internal/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadConfig()
	db := database.Connect()
	router := gin.Default()

	router.POST("/signup", handler.SignUp(db))
	router.POST("/login", handler.Login(db))

	auth := router.Group("/")
	auth.Use(middleware.AuthMiddleware())
	{
		auth.POST("/products", handler.CreateProduct(db))
		auth.GET("/products", handler.ListProducts(db))
		auth.POST("/orders", handler.CreateOrder(db))
	}

	router.Run(":8080")
}
