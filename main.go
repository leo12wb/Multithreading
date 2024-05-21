package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Address struct {
	CEP         string `json:"cep"`
	State       string `json:"state"`
	City        string `json:"city"`
	Neighborhood string `json:"neighborhood"`
	Street      string `json:"street"`
	Service     string `json:"service"`
}

func requestAPI(url string, ch chan Address, source string) {
	client := http.Client{Timeout: 1 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		ch <- Address{CEP: "", Service: source}
		return
	}
	defer resp.Body.Close()

	var address Address
	err = json.NewDecoder(resp.Body).Decode(&address)
	if err != nil {
		ch <- Address{CEP: "", Service: source}
		return
	}

	ch <- address
}

func main() {
	cep := "01153000"
	api1 := fmt.Sprintf("https://brasilapi.com.br/api/cep/v1/%s", cep)
	api2 := fmt.Sprintf("http://viacep.com.br/ws/%s/json/", cep)

	ch := make(chan Address, 2)

	go requestAPI(api1, ch, "BrasilAPI")
	go requestAPI(api2, ch, "ViaCEP")

	select {
	case address := <-ch:
		if address.CEP == "" {
			fmt.Printf("Erro ao obter dados do endereço da %s\n", address.Service)
			return
		}

		fmt.Printf("Dados do endereço obtidos pela API que respondeu mais rápido (%s):\n", address.Service)
		fmt.Printf("CEP: %s\n", address.CEP)
		fmt.Printf("Estado: %s\n", address.State)
		fmt.Printf("Cidade: %s\n", address.City)
		fmt.Printf("Bairro: %s\n", address.Neighborhood)
		fmt.Printf("Rua: %s\n", address.Street)
	case <-time.After(1 * time.Second):
		fmt.Println("Timeout ao tentar obter os dados do endereço.")
	}
}
