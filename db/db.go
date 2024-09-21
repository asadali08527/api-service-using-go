package db

import (
	"api-service/models"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

// InitDB initializes the database connection
func InitDB() *gorm.DB {
	dsn := "host=localhost user=postgres password=XXXXXX dbname=api_service port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Error connecting to the database: ", err)
	}
	// Migrate the schema
	err = db.AutoMigrate(&models.User{})
	if err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}
	DB = db
	return db
}
