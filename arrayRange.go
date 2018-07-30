//Edited with nano
package main
import (
	"fmt"
)
func main(){
	a := [...]float64{3,6,2,7,8}
	sum := float64(0)

	for i, v := range a{
		fmt.Printf("%d the element of a is %.2f\n",i ,v)
		sum += v
	}
	fmt.Println("The sum of all elements is:", sum)
}
