package main

import (
	"fmt"
	"time"
)

func main() {
	canal := escrever("caracoles")

	/*for mensagem := range canal {
		fmt.Println(mensagem)
	}*/

	for i := 0; i < 5; i++ {
		fmt.Println(<-canal)
	}
}

func escrever(texto string) <-chan string {
	canal := make(chan string)

	go func() {
		for {
			canal <- fmt.Sprintf("valor recebido: %s", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
