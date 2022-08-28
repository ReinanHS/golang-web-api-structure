package config

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/http/controller/user"
)

func AddRoutes(ctx context.Context, router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		prod := main.Group("user")
		{
			prod.GET("/", user.NewUserController(ctx).Index)
		}
	}

	return router
}
