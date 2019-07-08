package main

import "fmt"

func main(){
	for num:=0; num<10; num++{
		fmt.Println("Number is: ",num)

		if num==6 {
			break
		}
	}
}
