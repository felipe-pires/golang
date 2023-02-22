package main

import (
	"fmt"
	"modulo/15-testes/enderecos"
)

func main() {
	endereco := "avenida teste"

	resultado := enderecos.TipoDeEndereco(endereco)

	fmt.Println(resultado)
}
