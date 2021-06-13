package main

import (
	"fmt"
	"time"
)

func main() {
	i := 0

	for i < 2 {
		i++
		fmt.Println("add")
		time.Sleep(time.Second)
	}
	fmt.Println(i)

	for j := 0; j < 3; j++ {
		fmt.Println("add j:", j)
		time.Sleep(time.Second)
	}

	nomes := [3]string{"jao", "devi", "lucs"}
	//tb funciona pra slices
	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, letra := range "PALAVRA" {
		fmt.Println(letra)
		fmt.Println(string(letra))
	}

	usuario := map[string]string{
		"nome":      "jao",
		"sobrenome": "xilva",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	//for infinito
	for {
		fmt.Println("eterno...")
	}

}
