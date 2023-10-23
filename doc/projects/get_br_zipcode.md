# Get Address - Basic API

Help: [JSON-to-Go](https://mholt.github.io/json-to-go/)

```bash title="run command"
cd src/projectapi/
go run main.go
curl localhost:8080/get-address?zipcode=08295005
curl localhost:8080/get-address?zipcode=03087000
curl localhost:8080/get-address?zipcode=17509180
```

```go
--8<-- "src/projectapi/address/address_service.go"
```

```go
--8<-- "src/projectapi/main.go"
```

```bash title="output"
{"cep":"08295-005","logradouro":"Avenida Miguel Ignácio Curi","complemento":"","bairro":"Vila Carmosina","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"1004","ddd":"11","siafi":"7107"}
{"cep":"03087-000","logradouro":"Rua São Jorge","complemento":"","bairro":"Parque São Jorge","localidade":"São Paulo","uf":"SP","ibge":"3550308","gia":"1004","ddd":"11","siafi":"7107"}
{"cep":"17509-180","logradouro":"Avenida Vicente Ferreira","complemento":"até 470/471","bairro":"Marília","localidade":"Marília","uf":"SP","ibge":"3529005","gia":"4388","ddd":"14","siafi":"6681"}

```
