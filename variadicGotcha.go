package main
import (
	"fmt"
)
func change(s ...string){
	s[0] = "GO"
}

func main(){
	welcome := []string{"Welcome", "world"}

	change(welcome...)
	fmt.Println(welcome)
}
