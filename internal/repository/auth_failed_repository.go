package repository

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"gorm.io/gorm"
)

type AuthFailedRepository interface {
	Create(failed *entity.AuthFailed) error
}

type authFailedRepository struct {
	connection *gorm.DB
	ctx        context.Context
}

func NewAuthFailedRepository(ctx context.Context) AuthFailedRepository {
	return &authFailedRepository{
		connection: ctx.Value("db").(*gorm.DB),
		ctx:        ctx,
	}
}

func (c authFailedRepository) Create(failed *entity.AuthFailed) error {
	result := c.connection.Create(&failed)
	return result.Error
}
