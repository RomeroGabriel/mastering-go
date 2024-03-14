package main

import "fmt"

func main() {
	var nillMap map[string]int
	if nillMap == nil {
		fmt.Println("Not initalize map: ", nillMap)
		nillMap = make(map[string]int)
	}
	nillMap["key1"] = 1
	nillMap["key2"] = 2
	nillMap["key3"] = 3
	fmt.Println("After add elements: ", nillMap)

	value1 := nillMap["key1"]
	fmt.Println("Value from key1: ", value1)
	notExistKey := nillMap["keykkkk"]
	fmt.Println("When key doesn't exist, default value is returned : ", notExistKey)

	newMap := map[string]int{}
	newMap["Key1"] = 12
	fmt.Println("newMap: ", newMap)

	fullFillMap := map[string]int{
		"Key1": 12,
		"Key2": 22,
		"Key3": 32,
	}
	fmt.Println("fullFillMap: ", fullFillMap)
}
