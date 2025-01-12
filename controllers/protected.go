package controllers

import (
	"kenzoSec_task/database"
	"kenzoSec_task/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetTaskByID(c *gin.Context) {
	id := c.Param("id")
	var task models.Task
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}
	c.JSON(http.StatusOK, task)
}

// Implement CreateTask, UpdateTask, and DeleteTask similarly

func CreateTask(c *gin.Context) {
	var newTask models.Task

	// Bind incoming JSON to the newTask struct
	if err := c.ShouldBindJSON(&newTask); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Set default status if not provided
	if newTask.Status == "" {
		newTask.Status = "pending"
	}

	// Save the task to the database
	if err := database.DB.Create(&newTask).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create task"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "Task created successfully", "task": newTask})
}

func UpdateTask(c *gin.Context) {
	id := c.Param("id") // Retrieve the task ID from the URL parameter
	var task models.Task

	// Find the task by ID
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	var updateData models.Task
	// Bind incoming JSON to the updateData struct
	if err := c.ShouldBindJSON(&updateData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Update fields if they are provided
	if updateData.Title != "" {
		task.Title = updateData.Title
	}
	if updateData.Description != "" {
		task.Description = updateData.Description
	}
	if updateData.Status != "" {
		task.Status = updateData.Status
	}

	// Save the updated task to the database
	if err := database.DB.Save(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task updated successfully", "task": task})
}

func DeleteTask(c *gin.Context) {
	id := c.Param("id") // Retrieve the task ID from the URL parameter
	var task models.Task

	// Find the task by ID
	if err := database.DB.First(&task, id).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Task not found"})
		return
	}

	// Delete the task from the database
	if err := database.DB.Delete(&task).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete task"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Task deleted successfully"})
}
