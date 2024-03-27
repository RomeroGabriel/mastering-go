package main

import "fmt"

func main() {
	samples := []string{"hello", "apple_π!"}
outer:
	for _, sample := range samples {
		fmt.Println("Start sample: ", sample)
		for _, r := range sample {
			fmt.Println(string(r))
			if r == 'l' {
				continue outer
			}
		}
	}
}
