package router

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/http/controller"
	"github.com/reinanhs/golang-web-api-structure/internal/http/middleware"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// InitRouter configuration to add new router
func InitRouter(ctx context.Context, router *gin.Engine) *gin.Engine {

	main := router.Group("api/v1")
	{
		guest := main.Group("guest")
		{
			guest.POST("/register", controller.NewRegisteredUserController(ctx).Store)
			guest.POST("/login", controller.NewAuthenticatedSessionController(ctx).Store)
		}

		auth := main.Group("auth", middleware.Auth(ctx))
		{
			auth.GET("/me", controller.NewUserController(ctx).Me)
		}
	}

	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}
