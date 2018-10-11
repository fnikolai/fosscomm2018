package main

import "fmt"

func main() {
	// START OMIT
	fruitWeights := make(map[string]int) // HL
	fruitWeights["Apple"] = 45
	fruitWeights["Mango"] = 24
	fruitWeights["Orange"] = 34

	fmt.Printf("%#v\n", fruitWeights)
	// END OMIT
}
