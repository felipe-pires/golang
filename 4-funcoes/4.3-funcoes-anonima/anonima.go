package main

import "fmt"

func main() {
	func(n1, n2 int) {
		fmt.Println(n1 + n2)
	}(10, 20)
}
