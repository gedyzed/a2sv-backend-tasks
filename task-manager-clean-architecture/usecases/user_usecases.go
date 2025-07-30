package usecases

import (
	"context"
	"errors"
	"fmt"
	"task-manager-ca/domain"
)

type UserUsecase interface {

	Register (ctx context.Context, user *domain.User)(*domain.User, error)
	Login (ctx context.Context, user *domain.User)(string, error)
	PromoteAdmin (ctx context.Context, username string) error	
}

type OtherServices interface {
	HashPassword(password string)(string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
	GenerateToken(user *domain.User) (string, error)
}

type userUsecase struct {
	repo domain.UserRepository
	services OtherServices
}

func NewUserUsecase (repo domain.UserRepository, os OtherServices) UserUsecase {
	return &userUsecase{
		repo: repo,
		services: os, 
	}
}

func (u *userUsecase) Register (ctx context.Context, user *domain.User)(*domain.User, error){

	// checks if the username already taken
	existingUser, err := u.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		if errors.Is(err, domain.ErrWhileDecodingData){
			return nil, domain.ErrInternalServerError
		} else if errors.Is(err, domain.ErrWhileReadingData){
			return nil, domain.ErrInternalServerError 
		}
	}

	if existingUser != nil {
		return nil, domain.ErrUserAlreadyExists
	}

	
	// hash password
	user.Password, err = u.services.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// create the new user
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) Login (ctx context.Context, user *domain.User)(string, error){

	// get user using username
	existingUser, err := u.repo.GetByUsername(ctx, user.Username)
	fmt.Println(existingUser, err)
	if err != nil {
		return "", domain.ErrUsernameOrPassword
	}

	// check if the password matches
	err = u.services.CompareHashAndPassword(existingUser.Password, user.Password)
	if err != nil {
		return "", domain.ErrUsernameOrPassword
	}

	// generate and return token
	return u.services.GenerateToken(user)
}

func (u *userUsecase) PromoteAdmin (ctx context.Context, username string) error {
	return u.repo.Update(ctx, username)
}




