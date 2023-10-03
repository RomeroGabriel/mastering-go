package main

import "fmt"

type Team struct {
	TeamName string
}

type Person struct {
	Name string
	Age  int
	Team
	PersonTeam Team
}

func main() {
	person := Person{
		Name:       "Gabriel",
		Age:        100,
		Team:       Team{"Corinthians"},
		PersonTeam: Team{"Knicks"},
	}
	fmt.Printf("Person: %v\n", person)
	fmt.Printf("Person name: %v, Clube: %v\n", person.Name, person.Team)
	fmt.Printf("Person name: %v, Clube: %v\n", person.Name, person.Team.TeamName)
	fmt.Printf("Person name: %v, Clube: %v\n", person.Name, person.TeamName)
	fmt.Printf("Person name: %v, Clube: %v\n", person.Name, person.PersonTeam)
	fmt.Printf("Person name: %v, Clube: %v\n", person.Name, person.TeamName)
}
