package main

import "fmt"

//fortimente tipado

//global
var teste string

func main() {

	s := "felipe"
	fmt.Printf("%T %s \n", s, s)

	var s1 string
	s1 = "felipe"
	fmt.Println(s1)

	var (
		variavel  string = "teste"
		variavel2 string = "teste2"
	)

	variavel3, valivariavel4 := "teste3", "teste4"

	fmt.Println(variavel, variavel2, variavel3, valivariavel4)

	const contante string = "const"
	fmt.Println(contante)
}
