package main

import (
	"fmt"
	"regexp"
)

func verifyPassword(password1 string, password2 string) {
	matched, err := regexp.Match(password1, []byte(password2))
	fmt.Println(matched, err)
	if matched == true {
		fmt.Println("Password matched")
	} else {
		fmt.Println("Entered Password Mis-Match")
	}
}

// Let's test it
func main() {
	var password1 string
	var password2 string
	fmt.Println("set a password")
	fmt.Scanf("%s", &password1)
	fmt.Println("enter a password")
	fmt.Scanf("%s", &password2)

	verifyPassword(password1, password2)

}
