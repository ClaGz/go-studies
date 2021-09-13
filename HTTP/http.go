package main

import (
	"log"
	"net/http"
)

func main() {

	http.HandleFunc("/home", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Olá Mundo!"))
	})

	http.HandleFunc("/usuários", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Carregar página de usuários!"))
	})

	log.Fatal(http.ListenAndServe(":5000", nil))

}
