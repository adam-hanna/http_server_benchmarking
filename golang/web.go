package main

import (
	"fmt"
	"html/template"
	"net/http"
)

var myName string

type Page struct {
	Name string
}

var templates = template.Must(template.ParseFiles("template.html"))

func handler(w http.ResponseWriter, r *http.Request) {
	p := &Page{Name: myName}

	templates.ExecuteTemplate(w, "template.html", p)
}

func main() {
	fmt.Println("Listening on http://localhost:8080")
	myName = "World"

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
