package main

import (
	"fmt"
	"time"
)

// juntar 2 ou mais canais em um só

func main() {
	canal := multiplexar(escrever("texto 1"), escrever("texto 2"))

	for i := 0; i < 10; i++ {
		fmt.Println(<-canal)
	}
}

func multiplexar(canalEntrada1, canalEntrada2 <-chan string) <-chan string {
	canalSaida := make(chan string)

	go func() {
		for {
			select {
			case mensagem := <-canalEntrada1:
				canalSaida <- mensagem
			case mensagem := <-canalEntrada2:
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
			canal <- fmt.Sprintf("texto: %s ", texto)
			time.Sleep(time.Millisecond * 500)
		}
	}()

	return canal
}
