package main

import "fmt"

func main() {
	myMap := map[string]int{
		"key1": 1,
		"key2": 2,
		"key3": 3,
	}
	fmt.Println("Initial myMap: ", myMap)
	fmt.Println("key3 Value: ", myMap["key3"])
	myMap["key4"] = 4
	fmt.Println("Added key4: ", myMap)
	fmt.Println()

	if value, ok := myMap["key1"]; ok {
		fmt.Println("Key1 exists: ", value)
	}
	if value, ok := myMap["key5"]; ok {
		fmt.Println("Key5 exists: ", value)
	} else {
		fmt.Println("Key5 does not exist")
	}
	fmt.Println()
	delete(myMap, "key1")
	delete(myMap, "key5")
	fmt.Println("After delete key1 (exist) and key5 (not exist): ", myMap)
	fmt.Println("Len of myMap: ", len(myMap))

	for key, value := range myMap {
		fmt.Printf("Key: %s, Value: %d\n", key, value)
	}
	clear(myMap)
	fmt.Println(myMap)
}
