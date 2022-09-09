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

var logUserCreated = "usuário criado com sucesso"

//RegisteredUserController is a responsible to handle an incoming registration request.
type RegisteredUserController interface {
	Store(context *gin.Context)
}

type registeredUserController struct {
	userRepo        repository.UserRepository
	authSessionRepo repository.AuthSessionRepository
	userService     service.UserService
	authService     service.AuthService
	ctx             context.Context
}

//NewRegisteredUserController is creating anew instance of RegisteredUserController
func NewRegisteredUserController(ctx context.Context) RegisteredUserController {
	return &registeredUserController{
		userRepo:        repository.NewUserRepository(ctx),
		authSessionRepo: repository.NewAuthSessionRepository(ctx),
		userService:     service.NewUserService(ctx),
		authService:     service.NewAuthService(ctx),
		ctx:             ctx,
	}
}

// Store handle an incoming registration request
// @Summary     Registration of a new user
// @Param       user body dto.RegisterDto true "User JSON"
// @Description You will be able to create a user using this route
// @Tags        Attempt
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.User
// @Router      /guest/register [post]
func (c registeredUserController) Store(context *gin.Context) {
	params := dto.RegisterDto{}

	if data, err := params.BindingValidParams(context); err != nil {
		context.AbortWithStatusJSON(data.Code, data)
		return
	}

	user, err := c.userService.Store(params)
	if err != nil {
		request.ResponseDTO{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}.Abort(context)
		return
	}

	authSession, err := c.authService.GetAuthSession(context, &user)
	if err != nil {
		request.ResponseDTO{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}.Abort(context)
		return
	}

	_, err = c.authSessionRepo.Create(authSession)
	if err != nil {
		request.ResponseDTO{
			Message: err.Error(),
			Code:    http.StatusInternalServerError,
		}.Abort(context)
		return
	}

	code := http.StatusCreated
	context.JSON(code, request.ResponseDataDto{
		ResponseDTO: request.ResponseDTO{
			Message: logUserCreated,
			Code:    code,
		},
		Data: user,
	})
}
