package main

import "fmt"

func main() {
	// START OMIT
	var x []int // Empty slice
	fmt.Println(x, len(x), cap(x))

	x = append(x, 153, 146, 167, 170) // Resize and append // HL
	fmt.Println(x, len(x), cap(x))

	y := make([]int, 10, 20)       // Allocated slice (cap == 20)	// HL
	fmt.Println(y, len(y), cap(y)) // In use elements (len == 10)

	y = append(y, 153, 146, 167, 170) // Make a guess ...
	fmt.Println(y, len(y), cap(y))
	// END OMIT
}
