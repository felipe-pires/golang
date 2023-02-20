package main

import "fmt"

func main() {

	//so existe for

	//"while"
	i := 0
	for i < 10 {
		fmt.Println(i)
		i++
	}

	fmt.Println()
	fmt.Println()

	//"do while"
	j := 0
	for {
		fmt.Println(j)
		j++

		if j > 9 {
			break
		}
	}

	fmt.Println()
	fmt.Println()

	//range
	slice := []string{"A", "B", "C", "D", "E"}

	for indice, valor := range slice {
		fmt.Println(indice, valor)
	}

	fmt.Println()
	fmt.Println()
	//numero na tabela asc
	for indice, valor := range "PALAVRA" {
		fmt.Println(indice, valor)
	}
	fmt.Println()
	for indice, valor := range "PALAVRA" {
		fmt.Println(indice, string(valor))
	}

	fmt.Println()
	fmt.Println()

	usuario := map[string]string{
		"nome":      "felipe",
		"sobrenome": "pires",
	}

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
