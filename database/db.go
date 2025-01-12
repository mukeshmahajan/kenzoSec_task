package database

import (
	"fmt"
	"kenzoSec_task/models"
	"log"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init() {

	var dbErr error

	// Load .env file
	envErr := godotenv.Load()
	if envErr != nil {
		log.Println("No .env file found, using default settings")
	}

	// Read the database path from the environment variable
	dbPath := os.Getenv("DB_NAME")
	if dbPath == "" {
		dbPath = "task.db" // Default database path if not specified in .env
	}
	DB, dbErr = gorm.Open(sqlite.Open(dbPath), &gorm.Config{})
	if dbErr != nil {
		log.Fatal("Failed to connect to database:", dbErr)
	}
	fmt.Println("DB Connected")

	// Migrate the schema
	migrateErr := DB.AutoMigrate(&models.Task{})
	if migrateErr != nil {
		log.Fatal("Failed to migrate database:", migrateErr)
	}
}
