package routes

import (
	"kenzoSec_task/controllers"
	"kenzoSec_task/middleware"

	"github.com/gin-gonic/gin"
)

func ProtectedRoutes(router *gin.Engine) {
	protected := router.Group("/tasks")
	protected.Use(middleware.AuthMiddleware())
	protected.GET("/:id", controllers.GetTaskByID)
	protected.POST("/", controllers.CreateTask)
	protected.PUT("/:id", controllers.UpdateTask)
	protected.DELETE("/:id", controllers.DeleteTask)
}
