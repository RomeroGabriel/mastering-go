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
	fmt.Printf("Create object from an anonymous structs : %v\n", vector)
	fmt.Printf("X : %v\n", vector.X)
	fmt.Printf("Y : %v\n", vector.Y)
}
