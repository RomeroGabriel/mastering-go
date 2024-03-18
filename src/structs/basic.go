package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	var emptyPerson Person
	emptyPerson2 := Person{}

	fmt.Println(emptyPerson)
	fmt.Println(emptyPerson2)
	fmt.Println("Both empty structs are equal? ", emptyPerson == emptyPerson2)

	person := Person{
		Name: "Gabriel",
		Age:  100,
	}
	fmt.Printf("Created using field names: %v\n", person)

	person2 := Person{
		"Gabriel2",
		200,
	}
	fmt.Printf("Created using field order: %v\n", person2)
}
