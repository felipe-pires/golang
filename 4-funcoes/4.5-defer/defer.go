package main

import "fmt"

func funcao1() {
	fmt.Println("func 1")
}

func funcao2() {
	fmt.Println("func 2")
}

func main() {
	funcao1()
	funcao2()

	fmt.Println()
	fmt.Println()

	//adiar a execucao da funcao ate o ultimo momento possivel
	defer funcao1()
	funcao2()
}
