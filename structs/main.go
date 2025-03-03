package main

import (
	"fmt"
    "time"
)

type User struct {
    FirstName string
    LastName string
    Birthdate string
		createdOn time.Time
	}
    func (u User) out()  {
        fmt.Println("First Name: ", u.FirstName)}
func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")
	user := User{
		FirstName: firstName,
		LastName:  lastName,
		Birthdate: birthdate,
		createdOn: time.Now(),
	}
    user.out()

	
}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scan(&value)
	return value
}