package middleware

import (
	"context"
	"github.com/gin-gonic/gin"
	"github.com/reinanhs/golang-web-api-structure/internal/repository"
	"github.com/reinanhs/golang-web-api-structure/internal/service"
	"net/http"
)

func Auth(ctx context.Context) gin.HandlerFunc {
	return func(c *gin.Context) {
		const BearerSchema = "Bearer "
		header := c.GetHeader("Authorization")
		if header == "" {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		token := header[len(BearerSchema):]
		jwtService := service.NewJWTService(ctx)
		repo := repository.NewAuthAccessTokenRepository(ctx)

		if !jwtService.ValidateToken(token) {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		idFromToken, err := jwtService.GetIDFromToken(token)
		if err != nil {
			c.AbortWithStatus(http.StatusInternalServerError)
			return
		}

		auth, err := repo.GetUserByToken(idFromToken, token)
		if err != nil || auth.ID == 0 {
			c.AbortWithStatus(http.StatusNotFound)
			return
		}

		ctxRequest := context.WithValue(c.Request.Context(), "auth", auth)
		c.Request = c.Request.WithContext(ctxRequest)
		c.Next()
	}
}
