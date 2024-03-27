package main

import "fmt"

func main() {
	days := []string{"Monday", "Friday", "Tuesday", "Corinthians"}

	for _, day := range days {
		switch day {
		case "Monday":
			fmt.Println("Monday :(")
		case "Tuesday", "Wednesday", "Thursday":
			fmt.Println("Today is ", day)
		case "Friday":
			fmt.Println("IT'S FRIDAY FINALLY!")
		case "Saturday", "Sunday":
		default:
			fmt.Println("Is not a day")
		}
	}
	fmt.Println()
	fmt.Println("Now blank switch:")

	for _, word := range days {
		switch wordLen := len(word); {
		case wordLen <= 6:
			fmt.Println(word, "is a short day!")
		case wordLen > 6 && wordLen <= 10:
			fmt.Println(word, "is a long day!")
		default:
			fmt.Println(word, "is not a day!")
		}
	}
}
