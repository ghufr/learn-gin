package main

import (
	"github.com/ghufr/learn-gin/controllers"
	"github.com/ghufr/learn-gin/models"

	"github.com/gin-gonic/gin"
	// "github.com/swaggo/gin-swagger"
)

/*
	@title TODO list API
	@version 1.0

	@BasePath /api/v1
*/
func main() {
	router := gin.Default()

	models.ConnectDatabase()

	v1 := router.Group("/api/v1")
	{
		v1.GET("tasks", controllers.GetTasks)
		v1.GET("tasks/:id", controllers.GetTaskById)
		v1.POST("tasks", controllers.CreateTask)
		v1.PUT("tasks/:id", controllers.UpdateTask)
		v1.DELETE("tasks/:id", controllers.DeleteTask)
	}
	router.Run(":5000")
}