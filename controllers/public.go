package controllers

import (
	"kenzoSec_task/database"
	"kenzoSec_task/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//	func GetAllTasks(c *gin.Context) {
//		var tasks []models.Task
//		database.DB.Find(&tasks)
//		c.JSON(http.StatusOK, tasks)
//	}

func GetAllTasks(c *gin.Context) {
	var tasks []models.Task

	// Get query parameters
	pageStr := c.DefaultQuery("page", "1")    // Default to page 1
	limitStr := c.DefaultQuery("limit", "10") // Default to 10 items per page

	// Convert query parameters to integers
	page, err := strconv.Atoi(pageStr)
	if err != nil || page < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid page number"})
		return
	}

	limit, err := strconv.Atoi(limitStr)
	if err != nil || limit < 1 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid limit"})
		return
	}

	// Calculate offset
	offset := (page - 1) * limit

	// Query the database with pagination
	result := database.DB.Limit(limit).Offset(offset).Find(&tasks)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch tasks"})
		return
	}

	// Return paginated tasks
	c.JSON(http.StatusOK, gin.H{
		"page":  page,
		"limit": limit,
		"total": result.RowsAffected,
		"tasks": tasks,
	})
}
