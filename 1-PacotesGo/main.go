package main

import (
	"fmt"
	"modulo/1-PacotesGo/teste"

	"github.com/badoux/checkmail"
)

func main() {

	fmt.Println("hello word")
	teste.Teste()

	fmt.Println("valido: ", checkmail.ValidateFormat("felipe@teste.com"))
	fmt.Println("invalido: ", checkmail.ValidateFormat("felipeteste.com"))

}
