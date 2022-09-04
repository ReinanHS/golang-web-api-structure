package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/dto"
	"github.com/reinanhs/golang-web-api-structure/internal/repository"
	"net/http"
)

//RegisteredUserController is a ....
type RegisteredUserController interface {
	Store(context *gin.Context)
}

type registeredUserController struct {
	userRepository repository.UserRepository
	ctx            context.Context
}

//NewRegisteredUserController is creating anew instance of RegisteredUserController
func NewRegisteredUserController(ctx context.Context) RegisteredUserController {
	return &registeredUserController{
		userRepository: repository.NewUserRepository(ctx),
	}
}

// Store handle an incoming registration request
// @Summary     Registration of a new user
// @Param       user body dto.RegisterDto true "User JSON"
// @Description You will be able to create a user using this route
// @Tags        Auth
// @Accept      json
// @Produce     json
// @Success     200 {string} Helloworld
// @Router      /guest/register [post]
func (c registeredUserController) Store(context *gin.Context) {
	params := dto.RegisterDto{}

	if data, err := params.BindingValidParams(context); err != nil {
		context.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"data":    data,
			"message": err.Error(),
			"status":  http.StatusUnprocessableEntity,
		})
		return
	}

	context.JSON(200, gin.H{
		"data": params,
	})
}
