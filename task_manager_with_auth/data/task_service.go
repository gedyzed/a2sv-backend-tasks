package data

import (
	"context"
	"errors"
	"net/http"
	"task_manager_with_auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"github.com/gin-gonic/gin"
)

func GetTasks(c *gin.Context) (*[]models.Task, error) {
 
	var tasks []models.Task
	filter := bson.M{}
	cursor, err := TaskCollection.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(context.TODO())
	for cursor.Next(context.TODO()){
		var task models.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return &tasks, nil

}

func GetTaskById(c *gin.Context)(*models.Task, error){

	var task models.Task
	id := c.Param("id")
	filter := bson.M{"id":id}
	err := TaskCollection.FindOne(context.TODO(), filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func CreateTask(c *gin.Context)(*models.Task, int, error){

	var newTask models.Task
	if err := c.BindJSON(&newTask); err != nil {
		err = errors.New("Invalid Input Format")
		return nil, 400, err
	}

	filter := bson.M{"id":newTask.Id}
	result := TaskCollection.FindOne(context.TODO(), filter)
	err := result.Err()
	if err == nil  {
		err = errors.New("Task Id already exists")
		return nil, 400, err
	} 

	_, err = TaskCollection.InsertOne(context.TODO(), newTask)
	if err != nil {
		err = errors.New("Cannot Insert a Record")
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message" : "Cannot Insert a Record"})
		return nil, 500, err
	}

	return &newTask, 200, nil

}

func DeleteTask(c *gin.Context)(error){

	id := c.Param("id")
	filter := bson.M{"id": id}
	deleteResult, err := TaskCollection.DeleteOne(context.TODO(), filter)
	if err != nil || deleteResult.DeletedCount == 0 {
		return err
	}
	
	return nil
}

func UpdateTask(c *gin.Context)(*models.Task, int, error){

	id := c.Param("id")
	var updateTask models.Task
	if err := c.ShouldBindJSON(&updateTask); err != nil {
		err = errors.New("Invalid Input Format")
		return nil, 400, err
	}

	filter := bson.M{"id":id} 
	update := bson.M{}
	if updateTask.Description != "" {
		update["description"] = updateTask.Description
	}
	if updateTask.Status != "" {
		update["status"] = updateTask.Status
	}

	result, err := TaskCollection.UpdateOne(context.TODO(), filter, bson.M{"$set" :update})
	if err != nil {
		err = errors.New("Cannot Update Task")
		return nil, 500, err
	}
	if result.ModifiedCount == 0 {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message" : "Task Not Found"})
		err = errors.New("Task Not Found")
		return nil, 404, err
	}
	
	return &updateTask, 200, nil
}






