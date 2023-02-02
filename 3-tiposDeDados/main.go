package main

import (
	"errors"
	"fmt"
)

func main() {
	////////////////////////////////inteiros///////////////////////////////////
	var a int8
	var b int16
	var c int16
	var d int64
	var e int //utiliza a arquitetura da maquina como base
	var f uint

	a, b, c, d, e, f = 1, 123, 12345, 123456789, 1234545678, 123
	fmt.Println(a, b, c, d, e, f)

	// alias
	var varInt32 rune = 12345
	var varUint8 byte = 123

	fmt.Println(varInt32, varUint8)

	////////////////////////////////reais///////////////////////////////////

	var g float32 = 3.23
	var h float64 = 344.6754432545

	fmt.Println(g, h)

	/////////////////////////////string////////////////////////////////////

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

	z := `sasassa
	564545

	4545
	`
	fmt.Println(z)

	char := 'B' // numero na tabela asc
	char2 := string(char)
	fmt.Println(char, char2)

	///////////////////////////boolean/////////////////////////////////

	var u bool = false
	var v bool = true

	fmt.Println(u, v)

	//////////////////////////////error////////////////////////////
	var erro error = errors.New("not found")
	fmt.Println(erro)

	//valor zero
	var j string
	var k int
	var l float64
	var m bool
	var erro1 error
	fmt.Println(j, k, l, m, erro1)
}
