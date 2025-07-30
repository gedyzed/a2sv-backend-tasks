package infrastructure

import (
	"errors"
	"golang.org/x/crypto/bcrypt"
	"task_manager_ca/usecases"
)

type services struct{}

func NewServices() usecases.OtherServices {
	return &services{}

}


func(s *services) CompareHashAndPassword(hash string, password string) error {

	if bcrypt.CompareHashAndPassword([]byte(hash), []byte(password)) != nil {
		return errors.New("invalid password")
	}

	return nil

}

// Hash password
func (s *services) HashPassword(password string)(string, error){

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	password = string(hashedPassword)
	return password, nil
}
	

