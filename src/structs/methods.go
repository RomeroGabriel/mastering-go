package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func main() {
	circle := Circle{Radius: 5.0}
	area := circle.Area()
	fmt.Printf("Area: %v\n", area)
}
