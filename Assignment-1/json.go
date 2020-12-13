package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	bolType, _ := json.Marshal(false) //boolean Value
	fmt.Println(string(bolType))
	intType, _ := json.Marshal(10) // integer value
	fmt.Println(string(intType))
	fltType, _ := json.Marshal(3.14) //float value
	fmt.Println(string(fltType))
	strType, _ := json.Marshal("KloudOne") // string value
	fmt.Println(string(strType))
	slcA := []string{"golang", "python", "java"} //slice value
	slcB, _ := json.Marshal(slcA)
	fmt.Println(string(slcB))
	mapA := map[string]int{"golang": 1, "python": 2} //map value
	mapB, _ := json.Marshal(mapA)
	fmt.Println(string(mapB))
}
