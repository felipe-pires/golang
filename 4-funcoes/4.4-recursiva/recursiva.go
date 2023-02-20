package main

import "fmt"

func main() {

	posicao := 10
	for i := 1; i < posicao; i++ {
		fmt.Println(fibo(i))
	}

}

func fibo(posicao int) int {
	if posicao <= 1 {
		return posicao
	}

	return fibo(posicao-2) + fibo(posicao-1)
}
