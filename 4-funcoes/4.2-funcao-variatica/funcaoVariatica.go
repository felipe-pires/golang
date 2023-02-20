package main

import "fmt"

func main() {
	slice := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	fmt.Println(soma(slice...))

	fmt.Println(soma(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}

func soma(numeros ...int) int {
	soma := 0
	for _, valor := range numeros {
		soma += valor
	}

	return soma
}
