package repository

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"gorm.io/gorm"
)

//UserRepository is contract what userRepository can do to db
type UserRepository interface {
	All() []entity.User
	EmailIsInUsed(email string) bool
	UsernameIsInUsed(username string) bool
	Store(*entity.User) error
	RetrieveByCredentials(username string) *entity.User
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

func (r *userConnection) EmailIsInUsed(email string) bool {
	var count int64
	r.connection.Model(&entity.User{}).Where(&entity.User{Email: email}).Count(&count)

	return count != 0
}

func (r *userConnection) UsernameIsInUsed(username string) bool {
	var count int64
	r.connection.Model(&entity.User{}).Where(&entity.User{Username: username}).Count(&count)

	return count != 0
}

func (r *userConnection) Store(u *entity.User) error {
	result := r.connection.Create(u)
	return result.Error
}

func (r *userConnection) RetrieveByCredentials(username string) *entity.User {
	var user entity.User
	r.connection.
		Select([]string{"password", "username", "id"}).
		Where(&entity.User{Username: username}).
		First(&user)

	return &user
}
