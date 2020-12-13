package main

import "fmt"

func main() {

	//For
	sum := 0
	for i := 1; i < 5; i++ {
		sum += i
	}
	fmt.Println(sum) // 10 (1+2+3+4)

	//While Loop
	n := 1
	for n < 5 {
		n *= 2
	}
	fmt.Println(n) // 8 (1*2*2*2)

	//For-each
	strings := []string{"hello", "world"}
	for b, s := range strings {
		fmt.Println(b, s)
	}

	//Exit Loop
	d := 0
	for c := 1; c < 5; c++ {
		if c%2 != 0 { // skip odd numbers
			continue
		}
		d += c
	}
	fmt.Println(d) // 6 (2+4)

	//infinite loop
	//a := 0
	//for {
	//	a++ // repeated forever
	//}
	//fmt.Println(a) // never reached
}

