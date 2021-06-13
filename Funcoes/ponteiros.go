package main

import "fmt"

func inverterSinal(numero int) int {
	return numero * -1
}

func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	numero := 10
	numeroInvertido := inverterSinal(numero)
	fmt.Println(numero)
	fmt.Println(numeroInvertido)

	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
