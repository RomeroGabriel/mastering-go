package main

import "fmt"

func main() {
	empty_map := map[string]int{}
	fmt.Println(empty_map)
	empty_make := make(map[string]int)
	fmt.Println(empty_make)

	map_works_salary := map[string]int{"Gabriel": 100, "Renato Augusto": 3000}
	map_works_salary["Cassio"] = 3000
	gabriel_salary := map_works_salary["Gabriel"]
	fmt.Printf("Init map_works_salary: %v\n", map_works_salary)
	fmt.Printf("Gabriel salary: %v\n", gabriel_salary)

	value, exists := map_works_salary["Yuri Alberto"]
	if exists {
		fmt.Printf("Yuri Alberto salary: %v\n", value)
	} else {
		map_works_salary["Yuri Alberto"] = 3000
		yuri_salary := map_works_salary["Yuri Alberto"]
		fmt.Printf("Yuri Alberto salary: %v\n", yuri_salary)
	}

	delete(map_works_salary, "Gabriel")
	fmt.Printf("After delete Gabriel: %v\n\n", map_works_salary)

	for key, value := range map_works_salary {
		fmt.Printf("Player: %v, Salaray: %d\n", key, value)
	}
}
