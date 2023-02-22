package main

import "fmt"

func main() {
	canal := make(chan string, 2) // buffer especificar capacidade para o canal

	canal <- "teste 1"
	canal <- "teste 2"

	mensagem := <-canal
	mensagem2 := <-canal
	fmt.Println(mensagem, mensagem2)
}
