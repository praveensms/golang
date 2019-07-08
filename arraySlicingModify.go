package main
import (
	"fmt"
)
func main(){
	darr := [...]int{4,2,6,7,83,4,7,4}

	dslice := darr[2:5]

	fmt.Println("Array Before: ", darr)

	for i := range dslice{
		dslice[i]++
	}
	fmt.Println("Arrray after: ", darr)
}
