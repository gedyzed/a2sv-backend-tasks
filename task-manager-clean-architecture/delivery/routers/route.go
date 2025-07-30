package routers

import (
	"task-manager-ca/delivery/controllers"
	"task-manager-ca/infrastructure"
	"task-manager-ca/usecases"
	"task-manager-ca/repository"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func SetUp(db *mongo.Database, gin *gin.Engine){

	// all public routes
	publicRouter := gin.Group("")
	Register(db, publicRouter)
	Login(db, publicRouter)

	// users protected route
	userProtectedRoute := gin.Group("")
	userProtectedRoute.Use(
		infrastructure.AuthMiddleware(),
	)
	GetTasks(db, userProtectedRoute)
	GetTaskById(db, userProtectedRoute)
	CreateTask(db, userProtectedRoute)
	DeleteTask(db, userProtectedRoute)
	UpdateTask(db, userProtectedRoute)


	// admin protected routes
	adminProtectedRoute := gin.Group("")
	adminProtectedRoute.Use(
		infrastructure.AuthMiddleware(),
		infrastructure.AdminAuthMiddelware(),
	)
	PromoteAdmin(db, adminProtectedRoute)

}

func Register(db *mongo.Database, route *gin.RouterGroup){

	ur := repository.NewUserMongoRepo(db)
	os := infrastructure.NewServices()
	uc := usecases.NewUserUsecase(ur, os)
	rc := controllers.NewUserController(uc)

	route.POST("/register", rc.Register)
}

func Login(db *mongo.Database, route *gin.RouterGroup){

	ur := repository.NewUserMongoRepo(db)
	os := infrastructure.NewServices()
	uc := usecases.NewUserUsecase(ur, os)
	lc := controllers.NewUserController(uc)

	route.POST("/login", lc.Login)
}

func GetTasks(db *mongo.Database, route *gin.RouterGroup){
	
	tr := repository.NewTaskMongoRepo(db)
	tu := usecases.NewTaskUsecase(tr)
	tc := controllers.NewTaskController(tu)

	route.GET("/tasks", tc.GetTasks)

}

func GetTaskById(db *mongo.Database, route *gin.RouterGroup){

	tr := repository.NewTaskMongoRepo(db)
	tu := usecases.NewTaskUsecase(tr)
	tc := controllers.NewTaskController(tu)

	route.GET("/tasks/:id", tc.GetTaskById)


}

func CreateTask(db *mongo.Database, route *gin.RouterGroup){

	tr := repository.NewTaskMongoRepo(db)
	tu := usecases.NewTaskUsecase(tr)
	tc := controllers.NewTaskController(tu)

	route.POST("/create", tc.CreateTaskController)


}

func DeleteTask(db *mongo.Database, route *gin.RouterGroup){
	tr := repository.NewTaskMongoRepo(db)
	tu := usecases.NewTaskUsecase(tr)
	tc := controllers.NewTaskController(tu)

	route.DELETE("/delete/:id", tc.DeleteController)


}

func UpdateTask(db *mongo.Database, route *gin.RouterGroup){
	tr := repository.NewTaskMongoRepo(db)
	tu := usecases.NewTaskUsecase(tr)
	tc := controllers.NewTaskController(tu)

	route.PUT("/update/:id", tc.UpdateController)


}

func PromoteAdmin(db *mongo.Database, route *gin.RouterGroup){

	ur := repository.NewUserMongoRepo(db)
	os := infrastructure.NewServices()
	uc := usecases.NewUserUsecase(ur, os)
	ac := controllers.NewUserController(uc)

	route.POST("/promote-admin", ac.PromoteAdmin)
}

