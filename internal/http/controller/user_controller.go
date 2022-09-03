package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/repository"
)

//UserController is a ....
type UserController interface {
	Index(context *gin.Context)
}

type userController struct {
	userRepository repository.UserRepository
	ctx            context.Context
}

//NewUserController is creating anew instance of UserControlller
func NewUserController(ctx context.Context) UserController {
	return &userController{
		userRepository: repository.NewUserRepository(ctx),
	}
}

func (c userController) Index(context *gin.Context) {
	users := c.userRepository.All()

	context.JSON(200, gin.H{
		"data": users,
	})
}
