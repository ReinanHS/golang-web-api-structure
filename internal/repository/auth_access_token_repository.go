package repository

import (
	"context"
	"github.com/reinanhs/golang-web-api-structure/internal/dto"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"gorm.io/gorm"
)

type AuthAccessTokenRepository interface {
	Create(accessToken *entity.AuthAccessToken) error
	GetUserByToken(idFromToken int64, token string) (*dto.AuthDto, error)
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

func (c authAccessTokenRepository) GetUserByToken(idFromToken int64, token string) (*dto.AuthDto, error) {
	auth := &dto.AuthDto{}

	c.connection.Model(&entity.AuthAccessToken{}).
		Joins("inner join auth_sessions on auth_sessions.id = auth_access_tokens.auth_session_id").
		Joins("inner join users on users.id = auth_sessions.user_id").
		Where("auth_access_tokens.scopes = ?", token).
		Where("auth_access_tokens.is_revoked = ?", false).
		Where("auth_sessions.is_active = ?", true).
		Where("auth_sessions.id = ?", idFromToken).
		Select("users.id, users.name, users.username, users.avatar, users.email").
		Limit(1).
		Scan(&auth)

	return auth, nil
}
