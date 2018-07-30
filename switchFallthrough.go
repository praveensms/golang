package main

import "fmt"

func number() int{
	num := 15* 5
	return num
}

func main(){
	switch num := number();{
		case num < 50:
			fmt.Println("Number is less than 50")
			fallthrough
		case num < 100:
			fmt.Println("Number is less than 100")
			fallthrough
		case num < 200:
			fmt.Println("Number is less than 200")
			fallthrough
		default:
			fmt.Println("number is greater than 200")
	}
}
