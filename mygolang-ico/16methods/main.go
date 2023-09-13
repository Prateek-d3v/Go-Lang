package main

import "fmt"

func main() {
	fmt.Println("This is methods in Go language")

	// There is no such thing as Inheritance, superclasses or parent or child classes

	admin := User{"Prateek", "prateek@d3v.com", true, 24}
	fmt.Println(admin)
	fmt.Printf("Admin's details are: %+v \n", admin)
	fmt.Printf("Name is %v and Email is %v \n", admin.Name, admin.Email)
	admin.GetStatus()
	admin.NewMail()
	fmt.Printf("Name is %v and Email is %v \n", admin.Name, admin.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {
	fmt.Println("Is User Active:", u.Status)
}

func (u User) NewMail() {
	u.Email = "developer@d3v.com"
	fmt.Println("The new email is:", u.Email)
}
