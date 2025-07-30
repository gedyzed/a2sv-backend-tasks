package infrastructure

import (
	"task_manager_ca/domain"
	"github.com/dgrijalva/jwt-go"
)


func (s *services) GenerateToken(secret string, user *domain.User)(string, error){

	// Generate JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id":  user.UserID,
		"username": user.Username,
		"role":     user.Role,
	})

	jwtToken, err := token.SignedString(secret)
	if err != nil {
		return "", err
	}

	return jwtToken, err
}