package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/core/request"
)

type RegisterDto struct {
	Username  string `json:"username" form:"username" validate:"required"`
	Name      string `json:"name" form:"name" validate:"required"`
	Password  string `json:"password" form:"password" validate:"required"`
	CPassword string `json:"c_password" form:"c_password" validate:"required"`
	Email     string `json:"email" form:"email" validate:"required,email"`
}

func (params *RegisterDto) BindingValidParams(c *gin.Context) (request.ResponseErrorDto, error) {
	return request.DefaultGetValidParams(c, params)
}
