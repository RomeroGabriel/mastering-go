package main

import "fmt"

func main() {
	vector := struct {
		X int
		Y int
	}{
		X: 10,
		Y: 20,
	}
	fmt.Printf("Vector: %v\n", vector)
	fmt.Printf("X : %v\n", vector.X)
	fmt.Printf("Y : %v\n", vector.Y)

	var person struct {
		name string
		age  int
	}
	person.name = "Yuri Alberto"
	person.age = 123
	fmt.Printf("Person: %v\n", person)
	fmt.Printf("Name : %v\n", person.name)
	fmt.Printf("Age : %v\n", person.age)

}
