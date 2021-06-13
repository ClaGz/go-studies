package main

import "fmt"

//podemos ter mais parametros
// mas so podemos ter um parametro variático
// tambem ele precisa sempre ser o ultimo parametro
func soma(numeros ...int) (total int) {
	total = 0
	for _, numero := range numeros {
		total += numero
	}
	return
}

func main() {
	somaTotal := soma(1, 2, 3, 4, 1, 1, 1, 2, 2, 3, 1, 12, 3123, 123, 123, 123, 123)
	fmt.Println(somaTotal)

	somaTotalZero := soma()
	fmt.Println(somaTotalZero)
}
