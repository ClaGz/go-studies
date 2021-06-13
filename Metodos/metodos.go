package main

import "fmt"

type usuario struct {
	nome  string
	idade uint
}

func (u usuario) salvar() {
	fmt.Printf("dentro do metodo salvar. %s \n", u.nome)
}

func (u usuario) maiorDeIdade() bool {
	return u.idade >= 18
}

//tipo um setter, né?
func (u *usuario) fazerNiver() {
	u.idade++
}

func main() {
	usuario1 := usuario{"usuario1", 20}
	fmt.Println(usuario1)
	usuario1.salvar()

	deMaior := usuario1.maiorDeIdade()
	fmt.Println(deMaior)

	usuario1.fazerNiver()
	fmt.Println(usuario1.idade)
}
