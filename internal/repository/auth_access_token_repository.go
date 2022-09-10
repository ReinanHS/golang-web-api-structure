package repository

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"gorm.io/gorm"
)

type AuthAccessTokenRepository interface {
	Create(accessToken *entity.AuthAccessToken) error
}

type authAccessTokenRepository struct {
	connection *gorm.DB
	ctx        context.Context
}

func NewAuthAccessTokenRepository(ctx context.Context) AuthAccessTokenRepository {
	return &authAccessTokenRepository{
		connection: ctx.Value("db").(*gorm.DB),
		ctx:        ctx,
	}
}

func (c authAccessTokenRepository) Create(accessToken *entity.AuthAccessToken) error {
	result := c.connection.Create(&accessToken)
	return result.Error
}
