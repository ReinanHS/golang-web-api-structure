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

var apiURL = "http://ip-api.com/json"

type ipApiDto struct {
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

type IpApiService interface {
	GetInfoByIP(ip string) (ipApiDto, error)
}

type ipApiService struct {
	ctx context.Context
}

//NewIpApiService is creates a new instance of IpApiService
func NewIpApiService(ctx context.Context) IpApiService {
	return &ipApiService{
		ctx: ctx,
	}
}

func (s *ipApiService) GetInfoByIP(ip string) (ipApiDto, error) {
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s", apiURL, ip), nil)
	res, _ := http.DefaultClient.Do(req)

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Fatal(err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return ipApiDto{}, errors.New("não foi possivel consultar o IP")
	}

	body, _ := ioutil.ReadAll(res.Body)

	bodyData := string(body)
	dataJSON := ipApiDto{}
	err := json.Unmarshal([]byte(bodyData), &dataJSON)

	if err != nil {
		return ipApiDto{}, err
	}

	return dataJSON, nil
}
