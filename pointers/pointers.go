package main
import "fmt"
func main() {
	age:=32
	ageP := &age
	fmt.Println("Address of age variable is", &age)
	fmt.Println("Address of ageP variable is", ageP)
         write(ageP)}
func write(ageP *int) {
	fmt.Println("Value of ageP is", *ageP)
}