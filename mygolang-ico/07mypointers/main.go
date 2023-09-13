package main

import "fmt"

func main() {
	fmt.Println("Welcome to a tutorial on pointers")

	// var ptr *int
	// fmt.Println("Value of pointer is ", ptr)

	myNumber := 34

	var ptr = &myNumber
	fmt.Println("Value of actual pointer is ", ptr)
	fmt.Println("Value of actual pointer is ", *ptr)

	*ptr = *ptr + 2
	fmt.Println("New Value of pointer is ", myNumber)
}
