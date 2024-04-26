package main

import "fmt"

func wrongUpdate(p *int) {
	num2 := 20
	p = &num2
}

func Update(p *int) {
	*p = 20
}

func main() {
	x := 10
	wrongUpdate(&x)
	fmt.Println(x)
	Update(&x)
	fmt.Println(x)
}
