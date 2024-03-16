package main

import "fmt"

func main() {
	data := map[string]int{
		"key1": 1,
		"key2": 2,
	}

	v, ok := data["key1"]
	fmt.Println(v, ok)

	v, ok = data["key2"]
	fmt.Println(v, ok)

	if v, ok := data["key3"]; !ok {
		fmt.Println("Key3 not exists in data: ", v, ok)
	}

	data["key3"] = 3
	v, ok = data["key3"]
	fmt.Println(v, ok)
}
