package main

import "fmt"

func main() {
	fmt.Println("This is Structs in Go language")

	// There is no such thing as Inheritance, superclasses or parent or child classes

	prateek := User{"Prateek", "prateek@d3v.com", true, 24}
	fmt.Println(prateek)
	fmt.Printf("Prateek details are: %+v \n", prateek)
	fmt.Printf("Name is %v and Email is %v \n", prateek.Name, prateek.Email)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
