package repository

import (
	"context"
	"task_manager_ca/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
)


type mongoTaskRepo struct{
	coll *mongo.Collection
}

func NewTaskMongoRepo(db *mongo.Database) domain.TaskRepository{
	return &mongoTaskRepo{
		coll:db.Collection("tasks"),
	}
	
}

func (r *mongoTaskRepo) GetByID(ctx context.Context, id string)(*domain.Task, error){
	
	var task domain.Task
	filter := bson.M{"id": id}
	err := r.coll.FindOne(ctx, filter).Decode(&task)
	if err != nil {
		return nil, err
	}

	return &task, nil
}

func (r *mongoTaskRepo) GetTasks(ctx context.Context)([]domain.Task, error){

	var tasks []domain.Task
	filter := bson.M{}
	cursor, err := r.coll.Find(ctx, filter)
	if err != nil {
		return nil, err
	}

	defer cursor.Close(ctx)
	for cursor.Next(ctx){
		var task domain.Task
		err := cursor.Decode(&task)
		if err != nil {
			return nil, err
		}

		tasks = append(tasks, task)
	}

	return tasks, nil

}

func (r *mongoTaskRepo) Create(ctx context.Context, task *domain.Task)(*domain.Task, error) {

	filter := bson.M{"id": task.Id}
	result := r.coll.FindOne(ctx, filter)
	err := result.Err()
	if err == nil  {
		return nil, err
	} 

	_, err = r.coll.InsertOne(ctx, task)
	if err != nil {
		return nil, err
	}

	return task, nil
}

func (r *mongoTaskRepo) Update(ctx context.Context, task *domain.Task) (*domain.Task, error) {

	filter := bson.M{"id": task.Id} 
	update := bson.M{}

	if task.Description != "" {
		update["description"] = task.Description
	}
	if task.Status != "" {
		update["status"] = task.Status
	}

	result, err := r.coll.UpdateOne(ctx, filter, bson.M{"$set" :update})
	if err != nil {
		return nil, err
	}
	if result.ModifiedCount == 0 {
		return nil, err
	}
	
	return task, nil
}

func (r *mongoTaskRepo) Delete(ctx context.Context, id string) error {

	filter := bson.M{"id": id}
	deleteResult, err := r.coll.DeleteOne(ctx, filter)
	if err != nil || deleteResult.DeletedCount == 0 {
		return err
	}
	
	return nil
}
