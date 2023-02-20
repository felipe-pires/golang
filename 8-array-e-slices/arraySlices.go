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
}
