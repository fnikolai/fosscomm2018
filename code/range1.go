package main

import "fmt"

func main() {
	// START OMIT
	fruits := []string{"Apple", "Mango", "Orange"}
	for index, value := range fruits { // HL
		fmt.Println(index, value)
	}
	// END OMIT
}
