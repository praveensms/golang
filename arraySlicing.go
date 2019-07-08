package main
import (
	"fmt"
)
func main(){
	a := [...]int{3,6,2,56,8,9}

	var b []int  = a[1:4]


	fmt.Println("Sliced array is: ", b)
}
