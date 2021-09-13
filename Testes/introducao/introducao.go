package main

import (
	"fmt"
	"testes-auto/enderecos"
)

func main() {
	tipoEndereco := enderecos.TipoDeEndereco("Rodovia dos Imigrantes")
	fmt.Println(tipoEndereco)
}
