package main

import "fmt"

func main() {
	x := make([]string, 0, 5)
	x = append(x, "a", "b", "c", "d")
	y := x[:2:2]
	z := x[2:4:4]
	fmt.Println("x:", x, cap(x), len(x))
	fmt.Println("y:", y, cap(y), len(y))
	fmt.Println("z:", z, cap(z), len(z))
	fmt.Println()
	y = append(y, "i", "j", "k")
	x = append(x, "x")
	z = append(z, "y")
	fmt.Println("x:", x, cap(x), len(x))
	fmt.Println("y:", y, cap(y), len(y))
	fmt.Println("z:", z, cap(z), len(z))
}
