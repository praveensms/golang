package main

import "fmt"

func main(){

  a := true
  b := false

  fmt.Println("a is:", a, "b is:",b)

  c := a&&b
  fmt.Println("c is:", c)

  d := a||b
  fmt.Println("d is:", d)
}
