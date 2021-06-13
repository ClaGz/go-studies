package main

import "fmt"

func closure() func() {
	texto := "Dentro da closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "sou o texto da main"
	fmt.Println(texto)

	funcao := closure()
	funcao()
}
