package usecases

import (
	"context"
	"task_manager_ca/domain"
)

type UserUsecase interface {

	Register (ctx context.Context, user *domain.User)(*domain.User, error)
	Login (ctx context.Context, user *domain.User, secret string)(string, error)
	PromoteAdmin (ctx context.Context, username string) error	
}


type OtherServices interface {
	HashPassword(password string)(string, error)
	CompareHashAndPassword(hashedPassword string, password string) error
	GenerateToken(secret string, user *domain.User) (string, error)
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
	user, err := u.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return nil, err
	}

	// hash password
	user.Password, err = u.services.HashPassword(user.Password)
	if err != nil {
		return nil, err
	}

	// create the new user
	return u.repo.Create(ctx, user)
}

func (u *userUsecase) Login (ctx context.Context, user *domain.User, secret string)(string, error){

	// get user using username
	existingUser, err := u.repo.GetByUsername(ctx, user.Username)
	if err != nil {
		return "", err
	}

	// check if the password matches
	err = u.services.CompareHashAndPassword(existingUser.Password, user.Password)
	if err != nil {
		return "", err
	}

	// generate and return token
	return u.services.GenerateToken(secret, user)
}

func (u *userUsecase) PromoteAdmin (ctx context.Context, username string) error {
	return u.repo.Update(ctx, username)
}


