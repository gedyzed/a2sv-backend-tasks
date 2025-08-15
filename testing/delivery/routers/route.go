package routers

import (
	"task-manager-test/delivery/controllers"
	"task-manager-test/infrastructure"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(
	r *gin.Engine,
	userController *controllers.UserController,
	taskController *controllers.TaskController,
) {
	// Public routes
	publicRouter := r.Group("")
	{
		publicRouter.POST("/register", userController.Register)
		publicRouter.POST("/login", userController.Login)
	}

	// User protected routes
	userProtected := r.Group("")
	userProtected.Use(infrastructure.AuthMiddleware())
	{
		userProtected.GET("/tasks", taskController.GetTasks)
		userProtected.GET("/tasks/:id", taskController.GetTaskById)
		userProtected.POST("/create", taskController.CreateTaskController)
		userProtected.DELETE("/delete/:id", taskController.DeleteController)
		userProtected.PUT("/update/:id", taskController.UpdateController)
	}

	// Admin protected routes
	adminProtected := r.Group("")
	adminProtected.Use(
		infrastructure.AuthMiddleware(),
		infrastructure.AdminAuthMiddelware(),
	)
	{
		adminProtected.POST("/promote-admin", userController.PromoteAdmin)
	}
}
