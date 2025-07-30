package data

import (
	"go.mongodb.org/mongo-driver/bson"
	"context"
	"errors"
	"net/http"
	"task_manager_with_auth/models"

	"golang.org/x/crypto/bcrypt"
	"github.com/dgrijalva/jwt-go"

)

var jwtSecret = []byte("your_jwt_secret")

func RegisterUser(user *models.User) (int, *models.User, error) {

	// Check for duplicate username
	filter := bson.M{"username": user.Username}
	result := UserCollection.FindOne(context.TODO(), filter)
	if result.Err() == nil {
		return http.StatusConflict, nil, errors.New("username already exists")
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	// Check if user collection is empty
	cursor, err := UserCollection.Find(context.TODO(), bson.M{})
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to check user collection")
	}
	defer cursor.Close(context.TODO())

	if !cursor.Next(context.TODO()) {
		user.Role = "admin"
	} else {
		user.Role = "regular"
	}

	// Insert the user
	_, err = UserCollection.InsertOne(context.TODO(), user)
	if err != nil {
		return http.StatusInternalServerError, nil, errors.New("failed to insert user")
	}

	return http.StatusOK, user, nil
}

func LoginUser(user *models.User) (int, string, error) {
	
	filter := bson.M{"username": user.Username}
	result := UserCollection.FindOne(context.TODO(), filter)

	var oldUser models.User
	err := result.Decode(&oldUser)
	if err != nil || bcrypt.CompareHashAndPassword([]byte(oldUser.Password), []byte(user.Password)) != nil {
		return http.StatusUnauthorized, "", errors.New("invalid username or password")
	}

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  oldUser.ID.Hex(),
		"username": oldUser.Username,
		"role":     oldUser.Role,
	})

	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		return http.StatusInternalServerError, "", errors.New("failed to generate token")
	}

	return http.StatusOK, jwtToken, nil
}
func PromoteAdmin(username string) (int, error) {

	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": "admin"}}

	res, err := UserCollection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return 500, errors.New("error while updating data")
	}
	if res.MatchedCount == 0 {
		return 400, errors.New("user not found")
	}
	if res.ModifiedCount == 0 {
		return 200, errors.New("user already has admin role")
	}

	return 200, nil
}

