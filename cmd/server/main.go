package main

import (
	"servcio-api/internal/config"
	"servcio-api/internal/database"
	"servcio-api/internal/handler"
	"servcio-api/internal/middleware"

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
		auth.POST("/services", handler.CreateService(db))
		auth.GET("/services", handler.ListServices(db))
		auth.POST("/categories", handler.CreateCategory(db))
		auth.GET("/categories", handler.ListServices(db))
	}

	router.Run(":8080")
}
