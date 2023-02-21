package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("goroutine 1")
	go escrever("goroutine 2")
	escrever("teste")
}

func escrever(txt string) {
	for {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
