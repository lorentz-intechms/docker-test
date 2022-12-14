package main
 
import (
	"fmt"
)
 
func main() {
	aNumbers := []int{0, 1, 2, 3}
       

	
	// myMap := make(map[int]*int)
 
	for _,v :=range aNumbers{
		if v==1 {
			v=100
		}
	}
	for k,v :=range aNumbers{
		fmt.Println("k:",k,"v:",v)
	}
}