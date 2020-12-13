package main

import "fmt"

func main() {
	var a, b, sum int
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&a)
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&b)
	sum = a + b
	fmt.Print("Sum : ", sum)
}
package main

import "fmt"
	var a, b,c int

	
func main() {
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&a)
	fmt.Print("Enter a positive integer : ")
	fmt.Scan(&b)

    if a%b == 0 {
        fmt.Println("even")
    } else {
        fmt.Println("odd")
    }
	
	
fmt.Println("Type a number ")
fmt.Scan(&c)
    if c%4 == 0 {
        fmt.Println("It is divisible by 4")
    }

    if c := 9; c < 0 {
        fmt.Println(num, "is negative")
    } else if num < 10 {
        fmt.Println(num, "has 1 digit")
    } else {
        fmt.Println(num, "has multiple digits")
    }
}