package repository

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	All() []entity.User
}

type userConnection struct {
	connection *gorm.DB
	ctx        context.Context
}

//NewUserRepository is creates a new instance of UserRepository
func NewUserRepository(ctx context.Context) UserRepository {
	return &userConnection{
		connection: ctx.Value("db").(*gorm.DB),
		ctx:        ctx,
	}
}

func (r *userConnection) All() []entity.User {
	var users []entity.User

	r.connection.Debug().Select([]string{
		"id",
		"name",
		"email",
		"created_at",
		"updated_at",
	}).Find(&users)

	return users
}
