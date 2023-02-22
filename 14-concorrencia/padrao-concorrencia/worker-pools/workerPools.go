package main

import (
	"fmt"
)

func main() {
	tarefas := make(chan int, 45)
	resultados := make(chan int, 45)

	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)
	go worker(tarefas, resultados)

	for i := 0; i < 45; i++ {
		tarefas <- i
	}
	close(tarefas)

	for i := 0; i < 45; i++ {
		resultado := <-resultados
		fmt.Println(resultado)
	}

}

// tarefas <-chan int apenas recebe dados
//resultados chan<- int - apenas envia dados

func worker(tarefas <-chan int, resultados chan<- int) {
	for tarefa := range tarefas {
		resultados <- fibo(tarefa)
	}
}

func fibo(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibo(posicao-2) + fibo(posicao-1)
}
