package main

import (
	"fmt"
	htmlTemplate "html/template"
	"os"
	textTemplate "text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	tmpl, err := textTemplate.New("temple-string").Parse("Hello, {{ .Name }}")
	if err != nil {
		panic(err)
	}

	data := struct {
		Name string
	}{
		Name: "Gabriel",
	}
	tmpl.Execute(os.Stdout, data)
	fmt.Println("\n************")
	persons := [4]Person{
		{Name: "Gabriel", Age: 15},
		{Name: "Cassio", Age: 40},
		{Name: "Yuri", Age: 20},
		{Name: "Lucas", Age: 17},
	}

	tmpl_html := htmlTemplate.Must(
		htmlTemplate.New("person.html").ParseFiles("person.html"))
	tmpl_html.Execute(os.Stdout, persons)
}
