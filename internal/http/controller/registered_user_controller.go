package controller

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/core/request"
	"github.com/reinanhs/golang-web-api-structure/internal/dto"
	"github.com/reinanhs/golang-web-api-structure/internal/repository"
	"github.com/reinanhs/golang-web-api-structure/internal/service"
	"net/http"
)

//RegisteredUserController is a ....
type RegisteredUserController interface {
	Store(context *gin.Context)
}

type registeredUserController struct {
	userRepository repository.UserRepository
	userService    service.UserService
	ctx            context.Context
}

//NewRegisteredUserController is creating anew instance of RegisteredUserController
func NewRegisteredUserController(ctx context.Context) RegisteredUserController {
	return &registeredUserController{
		userRepository: repository.NewUserRepository(ctx),
		userService:    service.NewUserService(ctx),
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
		context.AbortWithStatusJSON(data.Code, data)
		return
	}

	data, err := c.userService.Store(params)
	if err != nil {
		code := http.StatusBadRequest
		context.AbortWithStatusJSON(code, request.ResponseDTO{
			Message: err.Error(),
			Code:    code,
		})
		return
	}

	context.JSON(200, gin.H{
		"data": data,
	})
}
