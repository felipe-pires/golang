package main

import "fmt"

//executa antes da main
func init() {
	fmt.Println("init")
}

func main() {
	fmt.Println("main")
}
