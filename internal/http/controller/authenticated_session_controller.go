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

//AuthenticatedSessionController is a responsible to handle an incoming authentication request.
type AuthenticatedSessionController interface {
	Store(context *gin.Context)
}

type authenticatedSessionController struct {
	userRepository repository.UserRepository
	authService    service.AuthService
	ctx            context.Context
}

//NewAuthenticatedSessionController is creating anew instance of AuthenticatedSessionController
func NewAuthenticatedSessionController(ctx context.Context) AuthenticatedSessionController {
	return &authenticatedSessionController{
		userRepository: repository.NewUserRepository(ctx),
		authService:    service.NewAuthService(ctx),
	}
}

// Store handle an incoming authentication request.
// @Summary     Perform user authentication
// @Param       user body dto.LoginDto true "User login JSON"
// @Description You will be able to create a user using this route
// @Tags        Attempt
// @Accept      json
// @Produce     json
// @Success     200 {object} entity.User
// @Router      /guest/login [post]
func (c authenticatedSessionController) Store(context *gin.Context) {
	params := dto.LoginDto{}
	if data, err := params.BindingValidParams(context); err != nil {
		context.AbortWithStatusJSON(data.Code, data)
		return
	}

	user, err := c.authService.Attempt(params)
	if err != nil {
		request.ResponseDTO{
			Message: err.Error(),
			Code:    http.StatusUnauthorized,
		}.Abort(context)
		return
	}

	_, err = c.authService.CheckSession(context, user)
	if err != nil {
		request.ResponseDTO{
			Message: err.Error(),
			Code:    http.StatusBadRequest,
		}.Abort(context)
		return
	}

	context.JSON(http.StatusCreated, gin.H{
		"data": "",
		"ip":   context.ClientIP(),
	})
}
