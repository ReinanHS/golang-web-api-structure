package request

import (
	"github.com/gin-gonic/gin"
)

type ResponseDTO struct {
	Message string `json:"message"`
	Code    int    `json:"code"`
}

func (r ResponseDTO) Abort(context *gin.Context) {
	context.AbortWithStatusJSON(r.Code, r)
}
