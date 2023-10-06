package main

import "fmt"

type Circle struct {
	Radius float64
}

func (c Circle) NormalArea() float64 {
	result := 3.14 * c.Radius * c.Radius
	c.Radius = 100
	return result
}

func (c *Circle) PointerArea() float64 {
	result := 3.14 * c.Radius * c.Radius
	c.Radius = 100
	return result
}

func main() {
	circle := Circle{Radius: 5.0}
	area := circle.NormalArea()
	fmt.Printf("Area: %v\n", area)
	fmt.Printf("Circle change? %v\n", circle)
	area = circle.PointerArea()
	fmt.Printf("Area: %v\n", area)
	fmt.Printf("Circle change? %v\n", circle)
}
