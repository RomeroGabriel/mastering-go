package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	jsonData := []byte(`{"name":"Gabriel","age":25}`)
	var person Person
	fmt.Println("Using Unmarshal")
	err := json.Unmarshal(jsonData, &person)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Person:", person)
	fmt.Println("***")
	fmt.Println("Using NewDecoder")
	file, err := os.Open("src/http/person.json")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	var person2 Person
	err = json.NewDecoder(file).Decode(&person2)
	fmt.Println("Person2:", person2)
}
