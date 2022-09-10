package dto

import (
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/core/request"
)

type LoginDto struct {
	Username string `json:"username" form:"username" validate:"required,min=4,max=60"`
	Password string `json:"password" form:"password" validate:"required,min=8"`
}

func (params *LoginDto) BindingValidParams(c *gin.Context) (request.ResponseErrorDto, error) {
	return request.DefaultGetValidParams(c, params)
}
