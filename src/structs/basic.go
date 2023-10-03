package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{
		Name: "Gabriel",
		Age:  100,
	}
	fmt.Printf("Create object person, from Person struct: %v\n", person)
}
