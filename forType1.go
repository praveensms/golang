package main

import (
	"fmt"
)

func main(){
	num := 0

	for ; num<=10; {
		fmt.Println("The number is:", num)
		num += 2
	}
}
