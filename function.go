package main

import "fmt"

func Add(varA int, varB int) int{
	add := varA + varB

	return add
}

func main(){
	num1, num2 := 12,23

	sum := Add(num1, num2)

	fmt.Println("The sum  iS:", sum)
}
