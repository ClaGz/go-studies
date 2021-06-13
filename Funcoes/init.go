package main

import "fmt"

var n int

func init() {
	//podemos ter uma função init por arquivo
	fmt.Println("antes do main")
	n = 20
}

func main() {
	//podemos ter uma função main por pacote
	fmt.Println("main e n:", n)
}
