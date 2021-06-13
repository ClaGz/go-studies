package main

import "fmt"

func recupera() {
	if r := recover(); r != nil {
		fmt.Println("Recuperamos")
	}
}

func calcAprovacao(n1, n2 uint) string {
	defer recupera() //vam recuperar o panico
	if media := (n1 + n2) / 2; media > 6 {
		return "aprovado"
	} else if media < 6 {
		return "reprovado"
	}

	panic("Exatamente 6 eu entro em panico!")

}

func main() {
	fmt.Println((calcAprovacao(6, 6)))
	fmt.Println("pos calcMedia")
}
