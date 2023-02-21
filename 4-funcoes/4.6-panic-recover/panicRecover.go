package main

import "fmt"

func main() {
	fmt.Println(test_panic(5, 5))
	fmt.Println("Pos execucao")
}

func test_panic(n1, n2 float64) bool {
	defer test_recover()
	media := (n1 + n2) / 2

	if media > 5 {
		return true
	} else if media < 5 {
		return false
	}
	panic("excecao")
}

func test_recover() {
	if r := recover(); r != nil {
		fmt.Println("test recover")
	}
}
