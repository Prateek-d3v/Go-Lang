package main

import (
	"fmt"
	"sort"
)

func main() {
	fruitList := []string{"Apple", "Orange", "Banana"}
	fmt.Printf("Type of fruitList is %T\n", fruitList)

	fruitList = append(fruitList, "Peach", "Grapes")
	fmt.Println(fruitList)

	fruitList = append(fruitList[1:4])
	fmt.Println(fruitList)

	highScores := make([]int, 4)

	highScores[0] = 773
	highScores[1] = 543
	highScores[2] = 642
	highScores[3] = 865
	// highScores[4] = 909

	highScores = append(highScores, 444, 555)

	fmt.Println(highScores)

	sort.Ints(highScores)
	fmt.Println(highScores)
	fmt.Println(sort.IntsAreSorted(highScores)) // True & False

	// How to remove a value from slice based on index

	var courses = []string{"reactjs", "javascript", "python", "swift", "ruby"}
	fmt.Println(courses)

	var index int = 3
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
