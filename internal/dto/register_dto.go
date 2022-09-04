package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/core/request"
)

type RegisterDto struct {
	Username  string `json:"username" form:"username" validate:"required,min=4,max=60"`
	Name      string `json:"name" form:"name" validate:"required,min=4,max=90"`
	Password  string `json:"password" form:"password" validate:"required,min=8"`
	CPassword string `json:"c_password" form:"c_password" validate:"required,min=8,eqfield=Password"`
	Email     string `json:"email" form:"email" validate:"required,email"`
}

func (params *RegisterDto) BindingValidParams(c *gin.Context) (request.ResponseErrorDto, error) {
	return request.DefaultGetValidParams(c, params)
}
