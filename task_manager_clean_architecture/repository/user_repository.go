package repository

import (
	"context"
	"task_manager_ca/domain"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/bson"
) 

type mongoUserRepo struct {
	coll *mongo.Collection
}

func NewUserMongoRepo(db *mongo.Database) domain.UserRepository {
	return &mongoUserRepo{
		coll: db.Collection("users"),
	}
}

func (r *mongoUserRepo) Create(ctx context.Context, user *domain.User)(*domain.User, error){

	cursor, err := r.coll.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	if !cursor.Next(ctx) {
		user.Role = "admin"
	} else {
		user.Role = "regular"
	}

	// Insert the user
	_, err = r.coll.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	return user, nil

}

func (r *mongoUserRepo) Update(ctx context.Context, username string) error {
	
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	res, err := r.coll.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return err
	}
	if res.MatchedCount == 0 {
		return err
	}
	if res.ModifiedCount == 0 {
		return err
	}

	return nil
}

func (r *mongoUserRepo) GetByUsername(ctx context.Context, username string)(*domain.User, error){

	// Check for duplicate username
	filter := bson.M{"username": username}
	result := r.coll.FindOne(ctx, filter)
	if result.Err() == nil {
		return nil, result.Err()
	}

	var user *domain.User
	result.Decode(&user)
	return user, nil
}











