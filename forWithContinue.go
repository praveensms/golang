package main

import "fmt"

func main(){
	for num := 0; num<10; num++{
		if num == 6 {
			fmt.Println("This number is skipped.")
			continue
		}
		fmt.Println("The number is: ", num)
	}
}
