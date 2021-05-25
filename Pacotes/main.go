package main

import (
	"Pacotes/auxiliar"
	"fmt"

	"github.com/badoux/checkmail"
)

//go mod tidy = remove todos os pacotes externos nao utilizados na sua app
func main() {
	fmt.Println("TESTE")
	auxiliar.Escrever()

	erro := checkmail.ValidateFormat("devbook@gmail.com")

	fmt.Println(erro)
}
