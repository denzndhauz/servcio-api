package database

import (
	"fmt"
	"log"
	"os"

	"ecomprac/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

// Connect connects to the DB and auto-migrates models
func Connect() *gorm.DB {
	dbType := os.Getenv("DB_TYPE")
	var db *gorm.DB
	var err error

	if dbType == "postgres" {
		dsn := fmt.Sprintf(
			"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
			os.Getenv("DB_HOST"),
			os.Getenv("DB_USER"),
			os.Getenv("DB_PASSWORD"),
			os.Getenv("DB_NAME"),
			os.Getenv("DB_PORT"),
		)
		db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	} else {
		log.Fatal("Unsupported database type")
	}

	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
	}

	// Auto-migrate all models
	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Order{},
		&models.OrderItem{},
	)
	if err != nil {
		log.Fatal("Failed to migrate database: ", err)
	}

	fmt.Println("Database connected and migrated successfully!")
	return db
}
