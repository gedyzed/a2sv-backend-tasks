package router

import (
	"task_manager/controllers"

	"github.com/gin-gonic/gin"
)

//create and map routers with their controllers
func Routers()*gin.Engine{

	router := gin.Default()
	router.GET("/tasks", controllers.GetTasks)
	router.GET("/tasks/:id", controllers.GetTaskById)
	router.POST("/tasks", controllers.AddTask)
	router.DELETE("/tasks/:id", controllers.DeleteTask)
	router.PUT("tasks/:id", controllers.UpdateTask)

	return router

}
