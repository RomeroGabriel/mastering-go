package main

import (
	"database/sql"
	"fmt"

	_ "github.com/mattn/go-sqlite3"
)

func main() {

	dumbDb, err := sql.Open("sqlite3", "")
	if err != nil {
		panic(err)
	}

	productUseCase := NewProductUseCase(dumbDb)
	product, err := productUseCase.GetProduct(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(product.Name)

	userUseCase := NewUserUseCase(dumbDb)
	user, err := userUseCase.GetUser(1)
	if err != nil {
		panic(err)
	}
	fmt.Println(user.Name)
}
