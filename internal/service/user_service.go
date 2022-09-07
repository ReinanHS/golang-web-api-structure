package service

import (
	"context"
	"errors"
	"fmt"
	"github.com/drexedam/gravatar"
	"github.com/reinanhs/golang-web-api-structure/internal/dto"
	"github.com/reinanhs/golang-web-api-structure/internal/entity"
	"github.com/reinanhs/golang-web-api-structure/internal/helper"
	"github.com/reinanhs/golang-web-api-structure/internal/repository"
)

var (
	errorInvalidEmail    = "o e-mail já está sendo usado"
	errorInvalidUsername = "o username já está sendo usado"
)

type UserService interface {
	Store(dto dto.RegisterDto) (entity.User, error)
}

type userService struct {
	repository         repository.UserRepository
	credentialsService CommonCredentialsService
	ctx                context.Context
}

//NewUserService is creates a new instance of UserService
func NewUserService(ctx context.Context) UserService {
	return &userService{
		repository:         repository.NewUserRepository(ctx),
		credentialsService: NewCommonCredentialsService(ctx),
		ctx:                ctx,
	}
}

func (s *userService) Store(dto dto.RegisterDto) (entity.User, error) {
	if s.repository.EmailIsInUsed(dto.Email) {
		return entity.User{}, errors.New(errorInvalidEmail)
	}

	if s.repository.UsernameIsInUsed(dto.Username) {
		return entity.User{}, errors.New(errorInvalidUsername)
	}

	isValidPassword, err := s.credentialsService.ValidPasswordsCommonCredentials(dto.Password)
	if err != nil || !isValidPassword {
		return entity.User{}, err
	}

	generateFromPassword, err := helper.GenerateFromPassword(dto.Password, helper.GetParams())
	if err != nil {
		return entity.User{}, err
	}

	user := entity.User{
		Username: dto.Username,
		Name:     dto.Name,
		Avatar: gravatar.New(dto.Email).
			Size(200).
			DefaultURL(fmt.Sprintf("https://avatars.dicebear.com/api/initials/%s.png?size=200", dto.Username)).
			AvatarURL(),
		Password: generateFromPassword,
		Email:    dto.Email,
	}

	if err := s.repository.Store(&user); err != nil {
		return entity.User{}, errors.New(err.Error())
	}

	return user, nil
}
