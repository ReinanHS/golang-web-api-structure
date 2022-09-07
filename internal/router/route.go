package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/http/controller"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter configuration to add new router
func InitRouter(ctx context.Context, router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		prod := main.Group("guest")
		{
			prod.POST("/register", controller.NewRegisteredUserController(ctx).Store)
			prod.POST("/login", controller.NewAuthenticatedSessionController(ctx).Store)
		}
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
