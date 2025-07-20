package controllers

import (
	"context"
	"fmt"
	"net/http"
	"task_manager_with_mongodb/data"
	"task_manager_with_mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)


func GetTasks(c *gin.Context){

	var tasks []models.Task
	current, err := data.Collection.Find(context.TODO(), bson.D{}) 
	if err != nil{
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Cannot Read Data"})
		return
	}

	for current.Next(context.TODO()){

		var task models.Task
		err := current.Decode(&task)
		if err != nil {
			c.IndentedJSON(http.StatusInternalServerError, gin.H{"message":"Error While File Reading"})
			return
		}
		tasks = append(tasks, task)
	}
	c.IndentedJSON(http.StatusOK, tasks)
	
}

func GetTaskById(c *gin.Context){

	var task models.Task
	id := c.Param("id")
	filter := bson.M{"id":id}
	err := data.Collection.FindOne(context.TODO(), filter).Decode(&task) 
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message":"Task Not Found"})
		return
	} 
	c.IndentedJSON(http.StatusOK,task)
}

func AddTask(c *gin.Context){

	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Invalid Input"})
		return
	}
 
	result, err := data.Collection.InsertOne(context.TODO(), newTask)
	fmt.Println(result)
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Cannot Insert a Record"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{
		"message" : "successfully created", 
		"task" : newTask})
}

func DeleteTask(c *gin.Context){

	id := c.Param("id")
	filter := bson.M{"id": id}
	deleteResult, err := data.Collection.DeleteOne(context.TODO(), filter)
	fmt.Println(deleteResult)
	if err != nil || deleteResult.DeletedCount == 0 {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Cannot delete a record"})
		return
	}
		c.IndentedJSON(http.StatusOK, gin.H{"message" : "Task delete successfully"})
}

func UpdateTask(c *gin.Context){

	id := c.Param("id")
	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message" : "Invalid Input"})
		return
	}

	filter := bson.M{"id":id} 
	update := bson.M{}
	if updateTask.Description != "" {
		update["description"] = updateTask.Description
	}
	if updateTask.Status != "" {
		update["status"] = updateTask.Status
	}

	fmt.Println(update)
	fmt.Println(update["description"])

	result, err := data.Collection.UpdateOne(context.TODO(), filter, bson.M{"$set" :update})
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Cannot Update Task"})
		return
	}
	fmt.Println(result.MatchedCount, result.ModifiedCount)
	if result.ModifiedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Task Not Found"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"message" : "Updated Successfully"})
	
}






