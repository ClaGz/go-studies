package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	numero     uint
	logradouro string
}

func main() {
	fmt.Println("STRUCTS")

	var u usuario
	fmt.Println(u) // { 0}

	u.nome = "tchongo"
	u.idade = 10
	fmt.Println(u) // {tchongo 10}

	end := endereco{2, "ruch"}
	u2 := usuario{"deivide", 21, end}
	fmt.Println(u2)

	u3 := usuario{nome: "tio"}
	fmt.Println(u3)

}
