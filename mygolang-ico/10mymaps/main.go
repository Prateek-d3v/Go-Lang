package main

import "fmt"

func main() {
	languages := make(map[string]string)

	languages["py"] = "python"
	languages["rb"] = "ruby"
	languages["js"] = "javascript"

	fmt.Println("List of languages are: ", languages)
	fmt.Println("py shorts for", languages["py"])

	// To delete element from map
	delete(languages, "rb")
	fmt.Println("List of languages are: ", languages)

	// for loops are interesting in GO language
	for key, val := range languages {
		fmt.Printf("For key %v value is %v \n", key, val)

	}
}
