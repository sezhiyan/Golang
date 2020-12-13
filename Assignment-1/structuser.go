package main

import "fmt"

type user struct {
	name string
	age  int
}

func newPerson(name string) *user {

	p := user{name: name}
	p.age = 42
	return &p
}

func main() {

	fmt.Println(user{"Bob", 20})

	fmt.Println(user{name: "Alice", age: 30})

	fmt.Println(user{name: "Fred"})

	fmt.Println(&user{name: "Ann", age: 40})

	fmt.Println(newPerson("Jon"))

	s := user{name: "Sean", age: 50}
	fmt.Println(s.name)

	sp := &s
	fmt.Println(sp.age)

	sp.age = 51
	fmt.Println(sp.age)
}
