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
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

//UserService is ...
type UserService interface {
	Store(dto dto.RegisterDto) (entity.User, error)
	ValidPasswordsCommonCredentials(pass string) bool
}

type userService struct {
	repository repository.UserRepository
	ctx        context.Context
}

//NewUserService is creates a new instance of UserService
func NewUserService(ctx context.Context) UserService {
	return &userService{
		repository: repository.NewUserRepository(ctx),
		ctx:        ctx,
	}
}

func (s *userService) Store(dto dto.RegisterDto) (entity.User, error) {
	if s.repository.EmailIsInUsed(dto.Email) {
		return entity.User{}, errors.New("o e-mail já está sendo usado")
	}

	if s.repository.UsernameIsInUsed(dto.Username) {
		return entity.User{}, errors.New("o username já está sendo usado")
	}

	if s.ValidPasswordsCommonCredentials(dto.Password) {
		return entity.User{}, errors.New("sua senha está na lista negra, digite outra senha")
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

func (s *userService) ValidPasswordsCommonCredentials(pass string) bool {
	url := "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10-million-password-list-top-1000.txt"
	req, _ := http.NewRequest("GET", url, nil)
	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)
	body, _ := ioutil.ReadAll(res.Body)

	return strings.Contains(string(body), pass)
}
