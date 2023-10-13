package mathpackage

import "fmt"

func sum[T int | float64](a, b T) T {
	fmt.Println("	Accessing sum function. Private to mathpackage")
	return a + b
}

func SumInt(a, b int) int {
	return sum(a, b)
}

func SumFloat(a, b float64) float64 {
	return sum(a, b)
}
