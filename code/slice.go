package main

import "fmt"

func main() {
	// START OMIT
	var x []int
	fmt.Println(x, len(x), cap(x))

	y := []int{153, 146, 167, 170}
	fmt.Println(y, len(y), cap(y))

	z := make([]int, 10)
	fmt.Println(z, len(z), cap(z))

	k := make([]int, 10, 20)
	fmt.Println(k, len(k), cap(k))
	// END OMIT
}
