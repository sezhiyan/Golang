package main  
  
import (  
 "fmt"  
)  
  
func main() {  
  
 //create a slice with string literal  
 countries := []string{"india", "USA", "UK", "France"}  
 fmt.Println(countries)  
 // Create and initilize a Map with make function  
 countriesMap := make(map[int]string)  
 // Iterate Slice  
 for i := 0; i < len(countries); i++ {  
  countriesMap[i] = countries[i]  
 }  
 fmt.Println(countriesMap)  
} 