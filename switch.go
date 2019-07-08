package main

import (
	"fmt"
)

func main(){
	finger := 3

	switch finger{
		case 1:
			fmt.Println("Thumb finger..")
		case 2:
			fmt.Println("Index finger..")
		case 3:
			fmt.Println("Middle finger.")
		case 4:
			fmt.Println("Ring finger...")
		case 5:
			fmt.Println("Little finger.")
		default:
			fmt.Println("This is wonder, he have a sixth finger. Wow! nice!")
	}
}
