package main

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

var templates *template.Template

type usuario struct {
	Nome  string
	Email string
}

func main() {

	templates = template.Must(template.ParseGlob("*.html"))

	http.HandleFunc("/home", func(rw http.ResponseWriter, r *http.Request) {
		u := usuario{Nome: "ASJHAS", Email: "asuasas@hsuais.com"}
		templates.ExecuteTemplate(rw, "home.html", u)
	})

	fmt.Println("Escutando")
	log.Fatal(http.ListenAndServe(":5000", nil))
}
