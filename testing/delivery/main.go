package main

import (
	"log"

	"task-manager-test/delivery/routers"
	"task-manager-test/delivery/controllers"
	"task-manager-test/infrastructure"
	"task-manager-test/repository"
	"task-manager-test/usecases"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Load .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found or failed to load it")
	}

	// Connect to database
	db := infrastructure.DbInit()

	// Get collections
	usersColl := db.Collection("users")
	tasksColl := db.Collection("tasks")

	// Create repos
	userRepo := repository.NewUserMongoRepo(usersColl)
	taskRepo := repository.NewTaskMongoRepo(tasksColl)

	// Create services
	services := infrastructure.NewServices()

	// Create usecases
	userUC := usecases.NewUserUsecase(userRepo, services)
	taskUC := usecases.NewTaskUsecase(taskRepo)

	// Create controllers
	userController := controllers.NewUserController(userUC)
	taskController := controllers.NewTaskController(taskUC)

	// Init Gin
	r := gin.Default()

	// Setup routes with injected controllers
    routers.RegisterRoutes(r, userController, taskController)
    


	// Start server
	if err := r.Run("localhost:8080"); err != nil {
		log.Fatal(err)
	}
}
