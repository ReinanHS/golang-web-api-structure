package service

import (
	"context"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/mileusna/useragent"
	"github.com/reinanhs/golang-web-api-structure/internal/dto"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"github.com/reinanhs/golang-web-api-structure/internal/helper"
	"github.com/reinanhs/golang-web-api-structure/internal/repository"
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
	Auth(context *gin.Context, dto dto.LoginDto) (*entity.AuthAccessToken, error)
}

type authService struct {
	repoUser        repository.UserRepository
	repoAuthSession repository.AuthSessionRepository
	repoAuthFailed  repository.AuthFailedRepository
	ctx             context.Context
}

func NewAuthService(ctx context.Context) AuthService {
	return &authService{
		repoUser:        repository.NewUserRepository(ctx),
		repoAuthSession: repository.NewAuthSessionRepository(ctx),
		repoAuthFailed:  repository.NewAuthFailedRepository(ctx),
		ctx:             ctx,
	}
}

func (s *authService) Auth(context *gin.Context, dto dto.LoginDto) (*entity.AuthAccessToken, error) {
	authToken := &entity.AuthAccessToken{}

	user, err := s.Attempt(dto)
	if err != nil {
		session, _ := s.GetAuthSession(context, user)
		failed := &entity.AuthFailed{
			UserId:    user.ID,
			Device:    session.Device,
			UserAgent: session.UserAgent,
			IpAddress: session.IpAddress,
		}

		_ = s.repoAuthFailed.Create(failed)
		return authToken, err
	}

	_, err = s.CheckSession(context, user)
	if err != nil {
		return authToken, err
	}

	return authToken, nil
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
