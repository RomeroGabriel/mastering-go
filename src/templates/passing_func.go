package main

import (
	"html/template"
	"os"
	"strings"
)

type Person struct {
	Name string
	Age  int
}

func printName(s string) string {
	return "Dear " + s
}

func main() {
	t := template.New("person_func.html")
	t.Funcs(template.FuncMap{"upperString": strings.ToUpper})
	t.Funcs(template.FuncMap{"printName": printName})
	persons := [4]Person{
		{Name: "Gabriel", Age: 15},
		{Name: "Cassio", Age: 40},
		{Name: "Yuri", Age: 20},
		{Name: "Lucas", Age: 17},
	}
	t = template.Must(t.ParseFiles("person_func.html"))
	t.Execute(os.Stdout, persons)
}
