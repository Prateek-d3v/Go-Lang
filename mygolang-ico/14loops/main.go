package main

import "fmt"

func main() {
	fmt.Println("Welcome to loops in Golang")

	days := []string{"Sun", "Mon", "Tue", "Wed", "Thu", "Fri", "Sat"}

	fmt.Println(days)

	// for d := 0; d < len(days); d++ {
	// 	fmt.Println(days[d])
	// }

	// for i := range days {  // i is index
	// 	fmt.Println(days[i])
	// }

	for index, day := range days {
		fmt.Printf("index is %v and day is %v\n", index, day)
	}
}
