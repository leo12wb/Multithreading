package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type APIResponse interface{}

type ApiResponse struct {
	Api string `json:"Api"`
}

type ViaCEP struct {
	ApiResponse
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

type ApiCEP struct {
	ApiResponse
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

func main() {
	channelViaCEP := make(chan ViaCEP)
	channelApiCEP := make(chan ApiCEP)

	go GetViaCEP(channelViaCEP)
	go GetApiCEP(channelApiCEP)

	select {
	case resViaCEP := <-channelViaCEP:
		fmt.Printf("ViaCEP: %+v\n", resViaCEP)

	case resApiCEP := <-channelApiCEP:
		fmt.Printf("ApiCEP: %+v\n", resApiCEP)

	case <-time.After(time.Second):
		fmt.Printf("TimeOut")

	}

}

func GetViaCEP(chApi chan ViaCEP) {
	var viaCEP ViaCEP
	RequestAPI("https://viacep.com.br/ws/22621252/json/", &viaCEP)
	viaCEP.Api = "ViaCEP"
	chApi <- viaCEP
}

func GetApiCEP(chVia chan ApiCEP) {
	var apiCEP ApiCEP
	RequestAPI("https://cdn.apicep.com/file/apicep/22260-003.json", &apiCEP)
	apiCEP.Api = "ApiCEP"
	chVia <- apiCEP
}

func RequestAPI(url string, res APIResponse) error {
	req, err := http.Get(url)
	if err != nil {
		return err
	}
	defer req.Body.Close()
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return err
	}
	err = json.Unmarshal(body, res)
	if err != nil {
		return err
	}
	return nil

}