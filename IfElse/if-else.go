package main

import "fmt"

func main() {
	numero := 10

	if numero > 15 {
		fmt.Println("maior")
	} else {
		fmt.Println("menor ou igual")
	}

	//A variavel outronumero só existe no contexto do if/else
	if outroNumero := numero; outroNumero > 0 {
		fmt.Println("numero maior que zero")
	} else {
		fmt.Println("numero menor ou igual a zero")
	}

}
