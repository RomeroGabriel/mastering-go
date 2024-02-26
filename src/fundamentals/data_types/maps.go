package main

import "fmt"

func main() {
	var nil_map map[string]int
	if nil_map == nil {
		fmt.Println("Not initalize map: ", nil_map)
		nil_map = make(map[string]int)
	}
	nil_map["key1"] = 1
	nil_map["key2"] = 2
	nil_map["key3"] = 3
	fmt.Println("After add elements: ", nil_map)

	value1 := nil_map["key1"]
	fmt.Println("Value from key1: ", value1)
	notExistKey := nil_map["keykkkk"]
	fmt.Println("When key doesn't exist, default value is returned : ", notExistKey)
}
