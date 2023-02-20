package main

import "fmt"

func main() {

	// referencia de memoria

	var v int
	var ponteiro *int

	v = 100
	ponteiro = &v

	fmt.Println(v, ponteiro)

	fmt.Println(*ponteiro) //desreferenciacao

	v = 150

	fmt.Println(ponteiro, *ponteiro)
}
