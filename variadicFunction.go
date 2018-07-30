package main
import (
	"fmt"
)
func Find(num int,nums ...int) {
	fmt.Printf("Type of nums is %T \n", nums)

	found := false

	for i,v := range nums {
		if v == num {
			fmt.Println("Number found at index", i , "in ", nums)
			found = true
		}
	}
	if !found {
		fmt.Println(num,"not found in ", nums)
	}
	fmt.Println("\n")
}

func main(){
	Find(12, 1, 2,5,12,1235)
	Find(4, 5, 7, 8, 4)
	Find(123, 4,23,4,456,678,234)
	Find(23,345,546,56,56,3,23)
}
