package main

import "fmt"

func main() {
	greeter()
	fmt.Println("Welcome to function in Golang.")
	// greeterTwo()

	result := adder(3, 5)
	proRes, myMessage := proAdder(4, 5, 6, 7, 8, 9, 1)

	fmt.Println("Result is:", result)
	fmt.Println("Result of proAdder is:", proRes)
	fmt.Println("Message of proAdder is:", myMessage)
}

func greeter() {
	fmt.Println("Hello, from Golang.")

}

func adder(val1 int, val2 int) int {
	return val1 + val2

}

func proAdder(values ...int) (int, string) {
	total := 0

	for _, val := range values {
		total += val
	}
	return total, "Successfully added"
}
