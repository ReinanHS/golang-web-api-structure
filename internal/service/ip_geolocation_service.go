package service

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

const ApiURL = "http://ip-api.com/json"

var errorBadRequest = "não foi possivel consultar o IP"

type IPGeolocationDto struct {
	Status      string      `json:"status"`
	Country     string      `json:"country"`
	CountryCode string      `json:"countryCode"`
	Region      string      `json:"region"`
	RegionName  string      `json:"regionName"`
	City        string      `json:"city"`
	Zip         string      `json:"zip"`
	Lat         json.Number `json:"lat,omitempty"`
	Lon         json.Number `json:"lon,omitempty"`
	Timezone    string      `json:"timezone"`
	Isp         string      `json:"isp"`
	Org         string      `json:"org"`
	As          string      `json:"as"`
}

type IPGeolocationService interface {
	GetInfoByIP(ip string) (IPGeolocationDto, error)
}

type ipGeolocationService struct {
	ctx context.Context
}

//NewIPGeolocationService is creates a new instance of IPGeolocationService
func NewIPGeolocationService(ctx context.Context) IPGeolocationService {
	return &ipGeolocationService{
		ctx: ctx,
	}
}

func (s *ipGeolocationService) GetInfoByIP(ip string) (IPGeolocationDto, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", ApiURL, ip), nil)
	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return IPGeolocationDto{}, errors.New(errorBadRequest)
	}

	body, _ := ioutil.ReadAll(res.Body)

	bodyData := string(body)
	dataJSON := IPGeolocationDto{}
	err := json.Unmarshal([]byte(bodyData), &dataJSON)

	if err != nil {
		return IPGeolocationDto{}, err
	}

	return dataJSON, nil
}
