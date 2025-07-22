package router

import (
	"task_manager_with_auth/controllers"
	"task_manager_with_auth/middleware"

	"github.com/gin-gonic/gin"
)

func Routers() *gin.Engine {

		Router := gin.Default()
		
		// Public routes
		Router.POST("/register", controllers.Register)
		Router.POST("/login", controllers.Login)

		// Authenticated user routes
		Router.GET("/tasks", middleware.AuthMiddleware(), controllers.GetTasks)
		Router.GET("/tasks/:id", middleware.AuthMiddleware(), controllers.GetTaskById)

		// Admin-only routes 
		Router.POST("/tasks/create", 
			middleware.AuthMiddleware(), 
			middleware.AdminAuthMiddelware(), 
			controllers.CreateTask,
		)

		Router.PUT("/tasks/edit/:id", 
			middleware.AuthMiddleware(), 
			middleware.AdminAuthMiddelware(), 
			controllers.UpdateTask,
		)

		Router.DELETE("/tasks/delete/:id", 
			middleware.AuthMiddleware(), 
			middleware.AdminAuthMiddelware(), 
			controllers.DeleteTask,
		)

		Router.POST("/promote-admin", 
			middleware.AuthMiddleware(), 
			middleware.AdminAuthMiddelware(), 
			controllers.PromoteAdmin,
		)


	return Router
}
