package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"github.com/reinanhs/golang-web-api-structure/internal/helper"
	"github.com/reinanhs/golang-web-api-structure/internal/repository"
	"github.com/reinanhs/golang-web-api-structure/pkg/dto"
	"github.com/reinanhs/golang-web-api-structure/pkg/entity"
)

var (
	errorInvalidCredentials = "credenciais Invalidas"
	errorInvalidSession     = "este dispositivo não está autorizado"
)

type AuthService interface {
	Attempt(dto dto.LoginDto) (*entity.User, error)
	AttemptSession(authSession *entity.AuthSession) (bool, error)
	CheckSession(context *gin.Context, user *entity.User) (bool, error)
	GetAuthSession(context *gin.Context, user *entity.User) (*entity.AuthSession, error)
	Auth(context *gin.Context, dto dto.LoginDto) (*dto.AuthTokenDto, error)
}

type authService struct {
	repoUser        repository.UserRepository
	repoAuthSession repository.AuthSessionRepository
	repoAuthFailed  repository.AuthFailedRepository
	repoAccessToken repository.AuthAccessTokenRepository
	jwtService      JWTService
	ctx             context.Context
}

func NewAuthService(ctx context.Context) AuthService {
	return &authService{
		repoUser:        repository.NewUserRepository(ctx),
		repoAuthSession: repository.NewAuthSessionRepository(ctx),
		repoAuthFailed:  repository.NewAuthFailedRepository(ctx),
		repoAccessToken: repository.NewAuthAccessTokenRepository(ctx),
		jwtService:      NewJWTService(ctx),
		ctx:             ctx,
	}
}

func (s *authService) Auth(context *gin.Context, dtoLogin dto.LoginDto) (*dto.AuthTokenDto, error) {
	user, err := s.Attempt(dtoLogin)
	if err != nil {
		session, _ := s.GetAuthSession(context, user)
		failed := &entity.AuthFailed{
			UserId:    user.ID,
			Device:    session.Device,
			UserAgent: session.UserAgent,
			IpAddress: session.IpAddress,
		}

		_ = s.repoAuthFailed.Create(failed)
		return nil, err
	}

	_, err = s.CheckSession(context, user)
	if err != nil {
		return nil, err
	}

	session, _ := s.GetAuthSession(context, user)
	sessionDB, err := s.repoAuthSession.GetAuthSessionByDeviceId(session.UserId, session.DeviceId)
	token, err := s.jwtService.GenerateToken(sessionDB.ID)
	if err != nil {
		return nil, err
	}

	authToken := &entity.AuthAccessToken{
		AuthSessionId: sessionDB.ID,
		IsRevoked:     false,
		Scopes:        token,
	}

	if err := s.repoAccessToken.Create(authToken); err != nil {
		return nil, err
	}

	return &dto.AuthTokenDto{
		Token:     token,
		ExpiresIn: 2,
	}, nil
}

func (s *authService) Attempt(dto dto.LoginDto) (*entity.User, error) {
	user := s.repoUser.RetrieveByCredentials(dto.Username)
	hash, _ := helper.ComparePasswordAndHash(dto.Password, user.Password)
	if !hash {
		return user, errors.New(errorInvalidCredentials)
	}

	return user, nil
}

func (s *authService) GetAuthSession(context *gin.Context, user *entity.User) (*entity.AuthSession, error) {
	uag := context.GetHeader("User-Agent")
	ua := useragent.Parse(uag)

	ipAddress := "24.48.0.1" // context.ClientIP()

	dataResult, err := NewIPGeolocationService(context).GetInfoByIP(ipAddress)
	if err != nil {
		return nil, err
	}

	session := entity.AuthSession{}
	session.UserId = user.ID
	session.User = *user
	session.Latitude = string(dataResult.Lat)
	session.Longitude = string(dataResult.Lon)
	session.Location = fmt.Sprintf("%s / %s", dataResult.RegionName, dataResult.City)
	session.IpAddress = ipAddress
	session.UserAgent = uag
	session.Device = ua.Name

	data := []byte(fmt.Sprintf(
		"%s-%s-%s-%s-%s",
		ua.Name,
		ua.OS,
		ua.Device,
		ipAddress,
		dataResult.CountryCode,
	))
	session.DeviceId = fmt.Sprintf("%x", md5.Sum(data))

	return &session, nil
}

func (s *authService) AttemptSession(authSession *entity.AuthSession) (bool, error) {
	session, err := s.repoAuthSession.GetAuthSessionByDeviceId(authSession.UserId, authSession.DeviceId)
	if err != nil && err.Error() == "record not found" {
		_, _ = s.repoAuthSession.Create(authSession)
	}

	if !session.IsActive {
		return false, errors.New(errorInvalidSession)
	}

	return true, nil
}

func (s *authService) CheckSession(context *gin.Context, user *entity.User) (bool, error) {
	session, _ := s.GetAuthSession(context, user)
	result, err := s.AttemptSession(session)

	return result, err
}
