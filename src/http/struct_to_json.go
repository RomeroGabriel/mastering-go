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
	person := Person{Name: "Gabriel", Age: 25}
	jsonData, err := json.Marshal(person)
	fmt.Println("Using Marshal")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("JSON Data:", string(jsonData))
	fmt.Println("***")
	fmt.Println("Using NewEncoder")
	file, _ := os.Create("src/http/person.json")
	defer file.Close()
	err = json.NewEncoder(file).Encode(person)
}
