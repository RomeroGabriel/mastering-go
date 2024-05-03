package main

import "fmt"

type Team struct {
	Name     string
	Location string
}

func (t Team) Print() string {
	return fmt.Sprintf("Team Name: %s, Location: %s", t.Name, t.Location)
}

type Person struct {
	Name string
	Age  int
	Team
}

func (p Person) Print() string {
	return fmt.Sprintf("Person Name: %s, Age: %d, and Team: %s", p.Name, p.Age, p.Team.Name)
}

func main() {
	person := Person{
		Name: "Gabriel",
		Age:  20,
		Team: Team{Name: "Corinthians", Location: "São Paulo"},
	}

	fmt.Println(person.Print())
	fmt.Println(person.Team.Print()) // Using field Team to access Team's print struct method
	fmt.Println(person.Location)     // Using Location field DIRECTLY
}
