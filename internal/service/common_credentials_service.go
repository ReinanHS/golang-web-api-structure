package service

import (
	"context"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

const ApiUrl = "https://raw.githubusercontent.com/danielmiessler/SecLists/master/Passwords/Common-Credentials/10-million-password-list-top-1000.txt"

var (
	errorFaildValidPassword = "não é possível realizar a consulta de senha da lista negra"
	errorInvalidPassword    = "sua senha está na lista negra, digite outra senha"
)

type CommonCredentialsService interface {
	ValidPasswordsCommonCredentials(pass string) (bool, error)
}

type commonCredentialsService struct {
	ctx context.Context
}

func NewCommonCredentialsService(ctx context.Context) CommonCredentialsService {
	return &commonCredentialsService{
		ctx: ctx,
	}
}

func (s *commonCredentialsService) ValidPasswordsCommonCredentials(pass string) (bool, error) {
	req, _ := http.NewRequest("GET", ApiUrl, nil)
	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return false, errors.New(errorFaildValidPassword)
	}

	body, _ := ioutil.ReadAll(res.Body)
	result := strings.Contains(string(body), pass)

	if result {
		return result, errors.New(errorInvalidPassword)
	}

	return result, nil
}
