package routes

import (
	"kenzoSec_task/controllers"

	"github.com/gin-gonic/gin"
)

func PublicRoutes(router *gin.Engine) {
	public := router.Group("/public")
	public.GET("/tasks", controllers.GetAllTasks)
}
