package main

import (
	"kenzoSec_task/database"
	"kenzoSec_task/migrations"
	"kenzoSec_task/routes"
	"log"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the database
	database.Init()

	// Run migrations
	migrations.RunMigrations()

	// Initialize Gin router
	router := gin.Default()

	// Setup routes
	routes.PublicRoutes(router)
	routes.ProtectedRoutes(router)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("defaulting to port %s", port)
	}
	router.Run(":" + port)
}
