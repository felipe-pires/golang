package teste

import "fmt"

//primeira letra maiuscula - deixar func publica
func Teste() {
	fmt.Println("Teste")
	teste1()
}

//primeira letra minuscula - func visivel apenas no pacote
func escrever() {
	fmt.Println("escrever")
}
