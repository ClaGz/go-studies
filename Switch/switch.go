package main

import "fmt"

func diaDaSemana(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sábado-feira"
	default:
		return "Número inválido"
	}
}

func diaDaSemana2(numero int) string {
	var dia string

	switch {
	case numero == 1:
		dia = "Domingo"
		fallthrough //Faz com que, após executado esse case, o próximo case tb seja executado
	case numero == 2:
		dia = "Segunda-feira"
	case numero == 3:
		dia = "Terça-feira"
	case numero == 4:
		dia = "Quarta-feira"
	case numero == 5:
		dia = "Quinta-feira"
	case numero == 6:
		dia = "Sexta-feira"
	case numero == 7:
		dia = "Sábado-feira"
	default:
		dia = "Número inválido"
	}
	return dia
}

func main() {

	dia := diaDaSemana(2)

	fmt.Println(dia)

	dia = diaDaSemana(5)

	fmt.Println(dia)

	dia = diaDaSemana(200)

	fmt.Println(dia)

}
