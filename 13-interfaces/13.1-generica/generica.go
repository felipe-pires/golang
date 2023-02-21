package main

import "fmt"

func generica(i interface{}) {
	fmt.Println(i)
}

func main() {
	generica("string")
	generica(10)
	generica(true)
}
