package controllers

import (
	"task-manager-ca/domain"
	"task-manager-ca/usecases"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	userUsecase usecases.UserUsecase
}

func(uc *UserController) Register(c *gin.Context){

	var user *domain.User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.IndentedJSON(400, gin.H{"message": "invalid input format"})
		c.Abort()
		return
	}

	created, err := uc.userUsecase.Register(c, user)
	if err != nil {
		c.IndentedJSON(500, gin.H{"message": err.Error()})
	}

	c.IndentedJSON(200, gin.H{"user": created })
}

