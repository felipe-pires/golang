package main

import "fmt"

func main() {
	soma, subtracao := calculo(20, 10)
	fmt.Println(soma, subtracao)
}

func calculo(n1, n2 int) (soma int, subtracao int) {
	soma = n1 + n2
	subtracao = n1 - n2
	return
}
