package main
import(
	"fmt"
)
func Array(num [5] int){
	num[0] = 12
	fmt.Println("Inside function ", num)
}

func main(){
	num := [...]int{5,6,7,8,9}

	fmt.Println("Before passing to function ", num)

	Array(num)

	fmt.Println("After passing to function ", num)
}
