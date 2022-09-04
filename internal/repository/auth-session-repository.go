package repository

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"gorm.io/gorm"
)

//AuthSessionRepository is contract what auth_session can do to db
type AuthSessionRepository interface {
	Create(*entity.AuthSession) (*entity.AuthSession, error)
	Delete(*entity.AuthSession)
	GetAuthSessionById(id int) (*entity.AuthSession, error)
	GetAll() []entity.AuthSession
}

type authSessionConnection struct {
	connection *gorm.DB
	ctx        context.Context
}

//NewAuthSession is creates a new instance of AuthSessionRepository
func NewAuthSession(ctx context.Context) AuthSessionRepository {
	return &authSessionConnection{
		connection: ctx.Value("db").(*gorm.DB),
		ctx:        ctx,
	}
}

//Create is responsible for registering in the database
func (c authSessionConnection) Create(authSession *entity.AuthSession) (*entity.AuthSession, error) {
	result := c.connection.Create(&authSession)
	return authSession, result.Error
}

//Delete is responsible for deleting in the database
func (c authSessionConnection) Delete(authSession *entity.AuthSession) {
	c.connection.Delete(&authSession)
	return
}

//GetAuthSessionById is responsible to do a search by id
func (c authSessionConnection) GetAuthSessionById(id int) (*entity.AuthSession, error) {
	var authSession = entity.AuthSession{}

	result := c.connection.First(&authSession, "id = ?", id)
	return &authSession, result.Error
}

//GetAll is responsible to bring all records
func (c authSessionConnection) GetAll() []entity.AuthSession {
	var authSessions []entity.AuthSession

	c.connection.Find(&authSessions)
	return authSessions
}
