package main  
import "fmt"  
import "time"
import "math/rand"
func worker(random chan int) { 
  rand.Seed(time.Now().Unix())  // seeding do that random number will produced  
    j := rand.Intn(100) 
   fmt.Print("working...")  
   time.Sleep(time.Second)
	  
   fmt.Println(j)  
   random <- j  
}  
func main() {  
   random := make(chan int, 1)  
   go worker(random)  
   <-random  
} 