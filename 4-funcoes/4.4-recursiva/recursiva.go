package main

import "fmt"

func main() {
	fmt.Println(fibo(10))
}

func fibo(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibo(posicao-2) + fibo(posicao-1)
}
