package main

import (
	"fmt"
	"time"
)

func main() {
	canal := multiplexador(escrever("arroz"), escrever("feijao"))

	for i := 0; i < 6; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexador(canal1, canal2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canal1:
				canalSaida <- mensagem
			case mensagem := <-canal2:
				canalSaida <- mensagem
			}
		}
	}()

	return canalSaida
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
