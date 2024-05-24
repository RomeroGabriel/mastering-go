package main

import "fmt"

type Shape interface {
	Area() float64
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func printArea(s Shape) {
	area := s.Area()
	fmt.Printf("Area: %.2f\n", area)
}

func main() {
	c := Circle{Radius: 5.0}
	r := Rectangle{Width: 4.0, Height: 6.0}
	printArea(c)
	printArea(r)
}
