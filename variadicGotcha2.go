package main
import (
	"fmt"
)
func change(s ...string){
	s[0] = "Go"

	s = append(s, "Play Ground")
	fmt.Println(s)
}
func main(){
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
