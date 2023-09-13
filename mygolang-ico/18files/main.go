package main

import "fmt"

func main() {
	name := "Alice"
	age := 30

	formattedString := fmt.Sprintf("Name: %s, Age: %d", name, age)
	// formattedString now contains "Name: Alice, Age: 30"

	fmt.Println(formattedString) // Print the formatted string
}
