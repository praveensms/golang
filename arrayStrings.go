package main
import (
	"fmt"
)
func main(){

	a := [...]string{"USA", "INDIA", "CHINA", "PAK"}

	b := a

	b[0] = "Singapore"

	fmt.Println("A array is: ", a)
	fmt.Println("B array is: ", b)
}
