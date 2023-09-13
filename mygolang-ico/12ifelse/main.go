package main

import "fmt"

func main() {

	fmt.Println("This is if else statement in go language")

	// var loginCount int = 10
	loginCount := 10
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch out"
	} else {
		result = "Excatly 10 login users"
	}

	fmt.Println(result)

	// For handling web requests
	if num := 4; num < 10 {
		fmt.Println("The numer is less than 10")

	} else if num > 10 {
		fmt.Println("The number is greater than 10")
	}
}
