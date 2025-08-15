package infrastructure

import (
	"fmt"
	"os"
	"task-manager-test/domain"

	"github.com/dgrijalva/jwt-go"
)


func (s *services) GenerateToken(user *domain.User)(string, error){

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserID,
		"username": user.Username,
		"role":     user.Role,
	})

	
    jwtSecret := []byte(os.Getenv("JWT_SECRET"))
	jwtToken, err := token.SignedString(jwtSecret)
	if err != nil {
		fmt.Println(err.Error())
		return "", domain.ErrInternalServerError
	}


	return jwtToken, nil
}