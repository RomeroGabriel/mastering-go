package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type Address struct {
	ZipCode     string `json:"cep"`
	PublicPlace string `json:"logradouro"`
	Extra       string `json:"complemento"`
	District    string `json:"bairro"`
	City        string `json:"localidade"`
	State       string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func main() {
	zip_codes := []string{"08295005", "03087000", "17509180"}
	for _, code := range zip_codes {
		url := "https://viacep.com.br/ws/" + code + "/json"
		req, err := http.Get(url)
		if err != nil {
			fmt.Println(err)
		}
		defer req.Body.Close()
		res, _ := io.ReadAll(req.Body)
		var address Address
		json.Unmarshal(res, &address)
		fmt.Println(address)
	}
}
