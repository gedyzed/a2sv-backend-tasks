package controllers

import (
	"net/http"
	"task-manager-test/domain"
	"task-manager-test/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func NewUserController(uc usecases.UserUsecase) *UserController {
	return &UserController{userUsecase: uc}
}

func (uc *UserController) Register(c *gin.Context) {

	// extract context
	ctx := c.Request.Context()

	var user *domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid input format"})
		c.Abort()
		return
	}


	created, err := uc.userUsecase.Register(ctx, user)

	if err != nil {
		switch err.Error() {
		case domain.ErrUserAlreadyExists.Error():
			c.IndentedJSON(409, gin.H{"message": err.Error()})
		default:
			c.IndentedJSON(500, gin.H{"message": err.Error()})
		}

		c.Abort()
		return
	}

	c.IndentedJSON(200, gin.H{"user": created})
}

func (uc *UserController) Login(c *gin.Context) {

	ctx := c.Request.Context()
	var user *domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid input format"})
		c.Abort()
		return
	}

	token, err := uc.userUsecase.Login(ctx, user)
	if err != nil {
		switch err.Error() {
		case domain.ErrUsernameOrPassword.Error():
			c.IndentedJSON(401, gin.H{"message": domain.ErrUsernameOrPassword})
		case domain.ErrInternalServerError.Error():
			c.IndentedJSON(500, gin.H{"message": err.Error()})

		}

		c.Abort()
		return
	}

	c.IndentedJSON(200, gin.H{"token": token})
}

func (uc *UserController) PromoteAdmin(c *gin.Context) {

	ctx := c.Request.Context()

	var user *domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid input format"})
		c.Abort()
		return
	}

	if user.Username != "" {
		c.IndentedJSON(400, gin.H{"error": "not enough fields"})
	}

	err := uc.userUsecase.PromoteAdmin(ctx, user.Username)
	if err != nil {
		switch err.Error() {
		case domain.ErrTheUserIsAdminAlready.Error():
			c.IndentedJSON(http.StatusNotModified, gin.H{"error": err.Error()})
		default:
			c.IndentedJSON(500, gin.H{"error": err.Error()})
		}

		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "the user has been prometed to admin"})
}

type TaskController struct {
	taskUsecase usecases.TaskUsecase
}

func NewTaskController(tu usecases.TaskUsecase) *TaskController {
	return &TaskController{taskUsecase: tu}
}

func (tc *TaskController) GetTasks(c *gin.Context) {

	ctx := c.Request.Context()

	tasks, err := tc.taskUsecase.GetTasks(ctx)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"tasks": tasks})
}

func (tc *TaskController) GetTaskById(c *gin.Context) {

	ctx := c.Request.Context()

	id := c.Param("id")
	task, err := tc.taskUsecase.GetTaskById(ctx, id)
	if err != nil {
		switch err.Error() {
		case domain.ErrTaskNotFound.Error():
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.IndentedJSON(500, gin.H{"error": err.Error()})
		}

		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"task": task})
}

func (tc *TaskController) CreateTaskController(c *gin.Context) {
	ctx := c.Request.Context()

	var task *domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid input format"})
		c.Abort()
		return
	}

	createdTask, err := tc.taskUsecase.AddTask(ctx, task)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{"task": createdTask})
}

func (tc *TaskController) DeleteController(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(400, gin.H{"message": "task id is required"})
		c.Abort()
		return
	}

	err := tc.taskUsecase.DeleteTask(ctx, id)
	if err != nil {
		switch err.Error() {
		case domain.ErrTaskNotFound.Error():
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.IndentedJSON(500, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "task deleted successfully"})
}

func (tc *TaskController) UpdateController(c *gin.Context) {
	ctx := c.Request.Context()

	id := c.Param("id")
	if id == "" {
		c.IndentedJSON(400, gin.H{"message": "task id is required"})
		c.Abort()
		return
	}

	var task *domain.Task
	if err := c.ShouldBindJSON(&task); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid input format"})
		c.Abort()
		return
	}

	// Set the ID from the URL parameter
	task.Id = id

	updatedTask, err := tc.taskUsecase.UpdateTask(ctx, task)
	if err != nil {
		switch err.Error() {
		case domain.ErrTaskNotFound.Error():
			c.IndentedJSON(http.StatusNotFound, gin.H{"error": err.Error()})
		default:
			c.IndentedJSON(500, gin.H{"error": err.Error()})
		}
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"task": updatedTask})
}
