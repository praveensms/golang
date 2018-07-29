package main

import (
  "fmt"
)

func main(){
  //Complex number examples

  comp1 := complex(5, 7)
  comp2 := 8 + 27i

  cadd := comp1 + comp2

  fmt.Println("Complex number addition is:", cadd)
}
