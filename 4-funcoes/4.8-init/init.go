package main

import "fmt"

//executa antes da main
func init() { //uma por arquivo
	fmt.Println("init")
}

func main() { //uma por pacote
	fmt.Println("main")
}
