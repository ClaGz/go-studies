package main

import "fmt"

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	fmt.Println(array1)

	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]int{1, 2, 5, 7, 8}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6, 6, 8, 4, 3, 4, 334, 3, 34, 3, 23, 23, 23, 23, 2, 23, 234, 2}
	fmt.Println(array3)

	//marromenos uma fatia de um array
	slice := []int{10, 11, 12, 14, 15}
	fmt.Println(slice)

	slice = append(slice, 34293) //append sempre devolve um slice novo
	fmt.Println(slice)

	slice2 := array3[1:4] //o 4 é exclusivo, ou seja, nao é pegado
	fmt.Println(slice2)

	//se eu alterar o array3 o slice tb sofre a mudança, pois ele é um ponteiro pra um array
	array3[2] = 453453
	fmt.Println("mudou = ", slice2)

	//Array interno
	//make aloca memória para o nosso sistema
	fmt.Println("-----SlICES -----")
	slice3 := make([]int, 10, 11)
	fmt.Println("tamanho atual:", len(slice3))
	fmt.Println("capacidade atual: ", cap(slice3))
	fmt.Println(slice3)

	slice3 = append(slice3, 89)
	fmt.Println("tamanho atual:", len(slice3))
	fmt.Println("capacidade atual: ", cap(slice3))
	fmt.Println(slice3)

	fmt.Println("caso vc passe da capacidade total do seu slice, o golang vai criar um outro array interno e isso vai dobrar o tamanho do seu slice")

	slice3 = append(slice3, 1010)
	fmt.Println("tamanho atual:", len(slice3))
	fmt.Println("capacidade atual: ", cap(slice3))
	fmt.Println(slice3)

	fmt.Println("Caso vc crie um slice sem capacidade...")
	slice4 := make([]int, 5)
	fmt.Println("tamanho atual:", len(slice4))
	fmt.Println("capacidade atual: ", cap(slice4))
	fmt.Println(slice4)

}
