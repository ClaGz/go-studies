package main

import (
	"fmt"
	"time"
)

func main() {
	canal := make(chan string)
	go escrever("tome -", canal)

	for {
		//sempre que eu faço [<- canal] eu espero a execuçao disso, é como um await
		mensagem, aberto := <-canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	}

	canal2 := make(chan string)
	go escrever("tome 2 -", canal2)

	//sempre que eu faço isso eu espero a execuçao disso, é como um await
	for mensagem := range canal2 {
		fmt.Println(mensagem)
	}

	fmt.Println("ACABOU")
}

func escrever(texto string, canal chan string) {
	for i := 0; i < 5; i++ {
		canal <- texto
		time.Sleep(time.Second)
	}

	close(canal)
}
