package migrations

import (
	"kenzoSec_task/database"
	"kenzoSec_task/models"
	"log"
)

// RunMigrations applies database migrations
func RunMigrations() {
	log.Println("Running database migrations...")

	// Ensure DB is not nil
	if database.DB == nil {
		log.Fatal("Database connection is nil. Please check the init function.")
	}

	// AutoMigrate creates or updates the database schema based on the Task model
	err := database.DB.AutoMigrate(&models.Task{})
	if err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	log.Println("Database migrations completed successfully.")
}
