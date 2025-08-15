package domain

import (
	"errors"
)


var (
	
    ErrUserAlreadyExists = errors.New("username already exists")
	ErrWhileReadingData = errors.New("error while reading data")
	ErrWhileDecodingData = errors.New("error while decoding data")
	ErrInternalServerError = errors.New("internal server error")
	ErrUserNotFound = errors.New("user not found")
	ErrUsernameOrPassword = errors.New("incorrect username or password")
	ErrTheUserIsAdminAlready = errors.New("the user has admin privilege already")
	ErrTaskNotFound = errors.New("task not found")
	ErrNoChange = errors.New("no modification has been applied")
)
