package controllers

import (
	"fmt"
	"net/http"
	"task_manager/data"
	"task_manager/models"

	"github.com/gin-gonic/gin"
)
  
func GetTasks(c *gin.Context){
	c.IndentedJSON(http.StatusOK, data.Tasks)
}

func GetTaskById(c *gin.Context){

	id := c.Param("id")
	for _, task := range data.Tasks {
		if task.Id == id {
			c.IndentedJSON(http.StatusOK, gin.H{"task": task})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "task not found"})
}

func AddTask(c *gin.Context){

	var newTask models.Task

	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid input"})
		return
	}

	for _, task := range data.Tasks {
		if task.Id == newTask.Id {
			c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Id already exists"})
			return 
		}
	} 

	data.Tasks = append(data.Tasks, newTask)
	c.IndentedJSON(http.StatusOK, gin.H{"message" : "successfully created"})
}

func DeleteTask(c *gin.Context){

	id := c.Param("id")
	for i, task := range data.Tasks {
		if task.Id == id {
			data.Tasks = append(data.Tasks[:i], data.Tasks[i + 1:]...)
			c.IndentedJSON(http.StatusOK, data.Tasks)
			return 
		}
	} 

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Task Not Found"})

}

func UpdateTask(c *gin.Context){

	id := c.Param("id")
	fmt.Println(id)
	var updateTask models.Task

	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "invalid input"})
		return
	}

	for i, task := range data.Tasks {
		if task.Id == id {
			if updateTask.Description != "" {
				data.Tasks[i].Description = updateTask.Description
			} else if updateTask.Status != "" {
				data.Tasks[i].Status = updateTask.Status
			}
			c.IndentedJSON(http.StatusOK, data.Tasks[i])
			return 
		}
	} 

	c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Task Not Found"})

}






