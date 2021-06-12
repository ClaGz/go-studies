package main

import "fmt"

type testP struct {
	valor int
}

func main() {
	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)

	var var3 int = 100
	var ponteiro *int
	fmt.Println(var3, ponteiro)

	ponteiro = &var3
	fmt.Println(ponteiro)
	fmt.Println("pra ver valor do ponteiro ", *ponteiro)

	testP1 := testP{10}
	testP2 := testP1
	fmt.Println("testP2 antes = ", testP2)

	testP1.valor = 100
	fmt.Println("testP2 depois = ", testP2)

	//golang copia objetos. nao passa referencia

	var p2 *testP = &testP1
	fmt.Println(p2)
	fmt.Println(*p2)
}
