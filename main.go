package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Address struct {
	CEP          string `json:"cep"`
	Logradouro   string `json:"logradouro"`
	Bairro       string `json:"bairro"`
	Localidade   string `json:"localidade"`
	UF           string `json:"uf"`
	Complemento  string `json:"complemento"`
	Numero       string `json:"numero"`
	Erro         bool   `json:"erro"`
	ErrorMessage string `json:"message"`
}

func requestAPI(url string, ch chan Address) {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ch <- Address{Erro: true, ErrorMessage: err.Error()}
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		ch <- Address{Erro: true, ErrorMessage: err.Error()}
		return
	}

	var address Address
	err = json.Unmarshal(body, &address)
	if err != nil {
		ch <- Address{Erro: true, ErrorMessage: err.Error()}
		return
	}

	address.Erro = false
	ch <- address
}

func main() {
	cep := "01153000"
	api1 := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	api2 := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	ch := make(chan Address, 2)

	go requestAPI(api1, ch)
	go requestAPI(api2, ch)

	var address Address
	select {
	case address = <-ch:
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout ao tentar obter os dados do endereço.")
		return
	}

	if address.Erro {
		fmt.Printf("Erro ao obter dados do endereço: %s\n", address.ErrorMessage)
		return
	}

	if address.Logradouro != "" {
		fmt.Println("Dados do endereço obtidos pela primeira API:")
	} else {
		fmt.Println("Dados do endereço obtidos pela segunda API:")
	}

	fmt.Printf("CEP: %s\n", address.CEP)
	fmt.Printf("Logradouro: %s\n", address.Logradouro)
	fmt.Printf("Bairro: %s\n", address.Bairro)
	fmt.Printf("Localidade: %s\n", address.Localidade)
	fmt.Printf("UF: %s\n", address.UF)
	fmt.Printf("Complemento: %s\n", address.Complemento)
	fmt.Printf("Número: %s\n", address.Numero)
}