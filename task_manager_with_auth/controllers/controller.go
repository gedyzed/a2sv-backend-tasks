package controllers

import (
	"fmt"
	"net/http"
	"task_manager_with_auth/data"
	"task_manager_with_auth/models"

	"github.com/gin-gonic/gin"
)

// User Controllers
func Register(c *gin.Context) {

	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"error":"Invalid Input Format"})
		return
	}

	statuscode, user, err := data.RegisterUser(user)
	if err != nil {
		c.IndentedJSON(statuscode, gin.H{"error": err.Error()})
		return
	}
	c.IndentedJSON(statuscode, gin.H{"message": "User Inserted Successfully", "user": user})
}

func Login(c *gin.Context) {

	var user *models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(401, gin.H{"error": "Invalid input format"})
		return
	}

	statuscode, jwtToken, err := data.LoginUser(user)
	if err != nil {
		c.IndentedJSON(statuscode, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(200, gin.H{"message": "User Logged Successfully", "token": jwtToken})
}

func GetTasks(c *gin.Context){

	tasks, err := data.GetTasks(c)
	if err != nil {
		c.IndentedJSON(500, gin.H{"error": "Error while reading data"})
		return
	}

	c.IndentedJSON(200, gin.H{"tasks": tasks})
}

func GetTaskById(c *gin.Context){

	task, err := data.GetTaskById(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Error while file reading"})
		c.Abort()
		return
	}

	c.IndentedJSON(200, gin.H{"task": task})
}


// Admin Controllers
func PromoteAdmin(c *gin.Context){
	
	var user models.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(401, gin.H{"error": "Invalid Input format"})
		c.Abort()
		return
	}

	statusCode, err := data.PromoteAdmin(user.Username)
	if err != nil{
		fmt.Println(err)
		c.IndentedJSON(statusCode, gin.H{"message": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(statusCode, gin.H{"message": "User has been Promoted to admin"})

}

func CreateTask(c *gin.Context){

	newTask, statusCode, err := data.CreateTask(c)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"message" : err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusCreated, gin.H{
		"message" : "successfully created", 
		"task" : newTask})

}

func DeleteTask(c *gin.Context){

	err := data.DeleteTask(c)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Cannot delete a record"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message" : "Task delete successfully"})

}

func UpdateTask(c *gin.Context){

	updatedFields, statusCode, err := data.UpdateTask(c)
	if err != nil {
		c.IndentedJSON(statusCode, gin.H{"error": err.Error()})
		c.Abort()
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{
		"message" : "Updated Successfully",
		"updated_fields":updatedFields,
	})
}






