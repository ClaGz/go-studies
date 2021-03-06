package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	soma := somar(10, 20)

	fmt.Println(soma)

	var f = func(text string) string {
		fmt.Println(text)
		return text
	}

	resultado := f("teste")

	fmt.Println(resultado)

	resultadoSoma, resultadoSubtracao := calculosMatematicos(10, 15)
	resultadoSomaDois, _ := calculosMatematicos(20, 25)

	fmt.Println(resultadoSomaDois)
	fmt.Println(resultadoSoma, resultadoSubtracao)
}
