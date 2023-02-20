package main

import "fmt"

func main() {
	// if init

	if numero := 10; numero > 0 {
		fmt.Println("if init")
	}

	//switch

	num := 3

	switch num {
	case 1:
		fmt.Println("switch")
	case 2:
		fmt.Println("switch")
	default:
		fmt.Println("switch")
	}

	switch {
	case num == 1:
		fmt.Println("switch")
	case num == 2:
		fmt.Println("switch")
	default:
		fmt.Println("switch")
	}

}
