package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int64 = 10000000
	fmt.Println(numero)

	//uint -> apenas numeros positivos
	//u vem de unsigned
	var numero2 uint32 = 1000
	fmt.Println(numero2)

	//alias
	// INT32 == RUNE
	var numero3 rune = 1235545
	fmt.Println(numero3)

	//alias
	// UINT = BYTE
	var numero4 byte = 123
	fmt.Println(numero4)

	//Quando eu nao passo o tipo, ele pega da arquitetura da minha maquina
	numero5 := 7128
	fmt.Println(numero5)

	var numeroReal float32 = 123.45
	fmt.Println(numeroReal)
	var numeroReal2 float64 = 1231231231.45
	fmt.Println(numeroReal2)
	numeroReal3 := 1232.122
	fmt.Println(numeroReal3)

	char := 'B' //aspas simples cria um char, mas um char devolve sempre um numero refetente ao caractere na tabela ascii
	fmt.Println(char)

	var texto int //todo tipo tem seu valor default
	fmt.Println(texto)

	var booleano bool
	fmt.Println(booleano)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("deu ruim")
	fmt.Println(erro2)
}
