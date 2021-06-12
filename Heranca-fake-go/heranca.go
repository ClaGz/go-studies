package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	fmt.Println("heranca so que nao")

	p1 := pessoa{"jose", "jose", 100, 143}
	fmt.Println(p1)

	e1 := estudante{p1, "pedreiragem", "vida"}
	fmt.Println(e1.nome)
	fmt.Println(e1)
}
