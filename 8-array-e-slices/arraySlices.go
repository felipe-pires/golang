package main

import (
	"fmt"
	"reflect"
)

func main() {

	//array

	var array1 [5]int
	fmt.Println(array1)

	array2 := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5, 6} // tamanho é fixo com base na quantidade de items passados
	fmt.Println(array3)

	//slices - sem tamanho fixo

	var slice []int
	fmt.Println(slice)

	// retorna um novo slice com os novos itens
	slice = append(slice, 1, 2, 3, 4, 5)
	fmt.Println(slice)
	fmt.Printf("\n%T %v\n", slice, slice)

	fmt.Print(reflect.TypeOf(slice))
	fmt.Print(reflect.TypeOf(array3))
	fmt.Println()
	fmt.Println()

	//Arrays Internos

	// arrays sao a base dos slices - quando um novo elemento é adiconado
	//a um slice um array é descartado e criado um novo que copia o anterior e  adiciona o novo item

	slice2 := make([]int, 10, 15) //tipo ,tamanho inicial, capacidade
	fmt.Println(slice2)

	fmt.Println(len(slice2)) //func length
	fmt.Println(cap(slice2)) // capacidade atual

	//quando a capacidade é ultrapassada é criado um novo array interno com o dobro do anterior
	slice2 = append(slice2, 1, 2, 3, 4, 5, 6)
	fmt.Println(slice2)

	fmt.Println(len(slice2))
	fmt.Println(cap(slice2))

}
