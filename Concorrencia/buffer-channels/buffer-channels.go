package main

import "fmt"

//Esse codigo daria deadlock se a gente n usasse o buffer
// pois sempre que eu faço [<-] meu programa para a execução para o canal
func main() {
	canal := make(chan string, 2) //to dizendo que o meu canal tem um limite de 2 mensagens = buffer
	canal <- "ola"
	canal <- "ola2"
	//canal <- "ola3" -> essa linha vai dar deadlock

	mensagem := <-canal
	mensagem2 := <-canal

	fmt.Println(mensagem)
	fmt.Println(mensagem2)
}
