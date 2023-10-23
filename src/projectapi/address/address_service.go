package address

import (
	"encoding/json"
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

func GetAddress(zipcode string) (*Address, error) {
	url := "https://viacep.com.br/ws/" + zipcode + "/json"
	req, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer req.Body.Close()
	res, _ := io.ReadAll(req.Body)
	var address Address
	err = json.Unmarshal(res, &address)
	if err != nil {
		return nil, err
	}
	return &address, nil
}
