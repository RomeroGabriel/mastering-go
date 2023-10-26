package main

import (
	"net/http"
	"text/template"
)

type Person struct {
	Name string
	Age  int
}

func main() {
	base_template := template.Must(template.ParseFiles(
		"header.html",
		"person.html",
		"footer.html",
	))

	persons := [4]Person{
		{Name: "Gabriel", Age: 15},
		{Name: "Cassio", Age: 40},
		{Name: "Yuri", Age: 20},
		{Name: "Lucas", Age: 17},
	}

	http.HandleFunc("/", func(res http.ResponseWriter, request *http.Request) {
		err := base_template.ExecuteTemplate(res, "header.html", persons)
		if err != nil {
			panic(err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
