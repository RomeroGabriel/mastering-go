package main

import "fmt"

type LinesOfText [][]byte

func printByLine(text LinesOfText) {
	for i, line := range text {
		fmt.Printf("Line %d text: %s\n", i, line)
	}
	fmt.Println()
}

func addLine(text LinesOfText, line string) LinesOfText {
	text = append(text, []byte(line))
	return text
}

func main() {
	text := LinesOfText{
		[]byte("Now is the time"),
		[]byte("for all good gophers"),
		[]byte("to bring some fun to the party."),
	}
	printByLine(text)
	text[0] = []byte("CHANGED LINE")
	printByLine(text)
	text = addLine(text, "NEW LINE")
	printByLine(text)
}
