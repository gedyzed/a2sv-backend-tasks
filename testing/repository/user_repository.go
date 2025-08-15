package repository

import (
	"context"
	"errors"
	"task-manager-test/domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
) 

type MongoSingleResult interface {
    Decode(v interface{}) error
    Err() error
}

type mongoUserRepo struct {
	coll MongoCollection
}

func NewUserMongoRepo(coll MongoCollection) domain.UserRepository {
	return &mongoUserRepo{
		coll: coll,
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
		return domain.ErrUserNotFound
	}
	
	if res.ModifiedCount == 0 {
		return domain.ErrNoChange
	}

	return nil
}

func (r *mongoUserRepo) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	filter := bson.M{"username": username}
	result := r.coll.FindOne(ctx, filter)

	if err := result.Err(); err != nil {
		if errors.Is(err, mongo.ErrNoDocuments) {
			return nil, domain.ErrUserNotFound
		}
		return nil, domain.ErrWhileReadingData
	}

	
	var user domain.User
	if err := result.Decode(&user); err != nil {
		return nil, domain.ErrWhileDecodingData
	}

	return &user, nil
}












