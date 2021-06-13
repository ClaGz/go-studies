package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		return fmt.Sprintf("o valor recebido foi: %s", texto)
	}("toma essa ai")

	fmt.Println(retorno)

	func() {
		fmt.Println("segura essaai")
	}()
}
