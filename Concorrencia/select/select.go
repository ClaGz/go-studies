package main

import (
	"fmt"
	"time"
)

func main() {
	canal1, canal2 := make(chan string), make(chan string)
	go func() {
		for {
			time.Sleep(time.Millisecond * 500)
			canal1 <- "CANAL 1"
		}
	}()

	go func() {
		for {
			time.Sleep(time.Second * 2)
			canal2 <- "CANAL 2"
		}
	}()

	/*for {
		//A mensagem do canal 1 ta chegando atrasada, pois estamos esperando a entrega do canal 2
		mensagemCanal1 := <-canal1
		fmt.Println(mensagemCanal1)

		//A mensagem do canal 2 ta chegando atrasada, pois estamos esperando a entrega do canal 1
		mensagemCanal2 := <-canal2
		fmt.Println(mensagemCanal2)
	}*/

	for {
		select {
		case mensagemCanal1 := <-canal1:
			fmt.Println(mensagemCanal1)
		case mensagemCanal2 := <-canal2:
			fmt.Println(mensagemCanal2)
		}
	}
}
