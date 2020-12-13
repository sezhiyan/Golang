package main

import "fmt"

func main() {
	a := 2
	switch a {
	case 1:
		fmt.Println("Sum is 1")
		fallthrough
	case 2:
		fmt.Println("Sum is 2")
		fallthrough
	case 3:
		fmt.Println("Sum is 3")
		fallthrough
	default:
		fmt.Println("Printing default")
	}
}


