package main

import "fmt"

func main() {
	defer fmt.Println("World!")
	defer fmt.Println("One")
	defer fmt.Println("Two")
	fmt.Println("Hello")
	myDefer()
}

// Last In, First Out

// Hello
// 43210Two
// One
// World!

func myDefer() {
	for i := 0; i < 5; i++ {
		defer fmt.Print(i)
	}
}
