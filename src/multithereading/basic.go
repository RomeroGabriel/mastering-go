package main

import (
	"fmt"
	"time"
)

func print_nice(name string) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d from %s\n", i, name)
		time.Sleep(1 * time.Second)
	}
}

func main() {
	go print_nice("Corinthians")
	go print_nice("Knicks")

	go func() {
		for i := 0; i < 4; i++ {
			fmt.Println("Just an annoying print")
			// time.Sleep(1 * time.Second)
		}
	}()

}
