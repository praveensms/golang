//Edited in nano editor
package main
import (
	"fmt"
)
func main(){
	a := 120
	var b * int = &a
	fmt.Printf("B is type of %T\n", b)
	fmt.Println("B value is: ", *b)
}
