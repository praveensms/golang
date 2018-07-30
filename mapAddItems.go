package main
import (
	"fmt"
)
func main(){
	personSalary := make(map[string]int)
	personSalary["Praveen"] = 10000
	personSalary["Mahendra"]= 20000
	personSalary["Harsha"]  = 30000

	fmt.Println("Salary of Praveen is: ", personSalary["Praveen"])
}
