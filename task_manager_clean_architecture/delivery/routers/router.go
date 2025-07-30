package router

import (
	controller "task_manager_ca/delivery/controllers"
	infrastructure "task_manager_ca/infrastructure"
	"task_manager_ca/usecases"
	"task-manager-ca/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUp(db *mongo.Database, gin *gin.Engine){

	// all public routes
	publicRouter := gin.Group("")
	Register(db, publicRouter)
	// Login(db, publicRouter)

	// // users protected route
	// userProtectedRoute := gin.Group("")
	// secret := []byte("your_jwt_secret")
	// userProtectedRoute.Use(
	// 	infrastructure.AuthMiddleware(&secret),
	// )


	// admin protected routes
	// adminProtectedRoute := gin.Group("")
	// adminProtectedRoute.Use(
	// 	infrastructure.AuthMiddleware(&secret),
	// 	infrastructure.AdminAuthMiddelware(),
	// )

}

func Register(db *mongo.Database, route *gin.RouterGroup){

	ur := repository.NewUserMongoRepo(db)
	os := infrastructure.NewServices()
	rc := controller.UserController{
		userUsecase: usecases.NewUserUsecase(ur, os),
	}

	route.POST("/register", rc.Register)
}

// func Login(db *mongo.Database, route *gin.RouterGroup){


// }

// func GetTasks(db *mongo.Database, route *gin.RouterGroup){


// }

// func GetTaskById(db *mongo.Database, route *gin.RouterGroup){


// }

// func CreateTask(db *mongo.Database, route *gin.RouterGroup){


// }

// func DeleteTask(db *mongo.Database, route *gin.RouterGroup){


// }

// func UpdateTask(db *mongo.Database, route *gin.RouterGroup){


// }

