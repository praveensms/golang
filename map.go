package main
import (
	"fmt"
)
func main(){
	var personSalary map[string]int
	if personSalary == nil {
		fmt.Println("personSalary is nill, going to create new one..")
		personSalary = make(map[string]int)
	}
	fmt.Println(personSalary)
}
