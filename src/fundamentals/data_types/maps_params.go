package main

import "fmt"

func mapParam(data map[string]string) {
	fmt.Printf("Memory allocated for data: %p\n", &data)
	data["key3"] = "Liverpool"
}

func main() {
	myMap := make(map[string]string)
	myMap["key1"] = "Corinthians"
	myMap["key2"] = "Lakers"

	fmt.Println("Inital mySlice: ", myMap)
	fmt.Printf("Memory allocated for mySlice: %p\n", &myMap)
	mapParam(myMap)
	fmt.Println("After call func Slice: ", myMap)
}
