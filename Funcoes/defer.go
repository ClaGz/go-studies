package main

import "fmt"

func func1() {
	fmt.Println("func 1")
}

func func2() {
	fmt.Println("func 2")
}

func main() {
	//func1 vai ser a ultima execução antes de sairmos da função
	// podemos dizer, logo antes do retorno
	defer func1()

	func2()
}
