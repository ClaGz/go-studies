package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("olha")
	escrever("olha tambem")
}

func escrever(texto string) {
	for {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
