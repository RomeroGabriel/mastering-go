package main

import "fmt"

func main() {
	wrongWay := make([]int, 5)
	wrongWay = append(wrongWay, 10)
	wrongWay[1] = 20
	fmt.Println("Notice that the value 10 is placed in the end: ", wrongWay)
	fmt.Println()
	fmt.Println("Creating slice is len 0 but with cap--->")
	rightWay := make([]int, 0, 5)
	rightWay = append(rightWay, 10)
	fmt.Println("Notice that now value 10 is placed in the begin: ", rightWay, len(rightWay), cap(rightWay))
	rightWay = append(rightWay, 1, 2, 3, 4)
	fmt.Println(rightWay)
}
