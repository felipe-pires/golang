package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var waitGroup sync.WaitGroup

	waitGroup.Add(2)

	go func() {
		escrever("goroutine 1")
		waitGroup.Done()
	}()

	go func() {
		escrever("goroutine 2")
		waitGroup.Done()
	}()

	waitGroup.Wait()

}

func escrever(txt string) {
	for i := 0; i < 5; i++ {
		fmt.Println(txt)
		time.Sleep(time.Second)
	}
}
