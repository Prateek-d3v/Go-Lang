package main

import "fmt"

func main() {
	fmt.Println("Welcome to array in golang")

	var fruitList [5]string

	fruitList[0] = "Apple"
	fruitList[1] = "Tomato"
	fruitList[3] = "Peach"
	fruitList[4] = "Bacon"

	fmt.Println("fruit list: ", fruitList)
	fmt.Println("fruit list: ", len(fruitList))

	var vegList = [3]string{"Bacon", "Beans", "Potato"}
	fmt.Println("Vegy list: ", vegList)
}
