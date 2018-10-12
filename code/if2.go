package main

import "fmt"

func main() {
	// START OMIT
	b, c := 10, 31
	if mean := b + c; mean < 42 {
		fmt.Println("if:", mean)
	} else {
		fmt.Println("else:", mean-42)
	}
	fmt.Println("outside:", mean)
	// END OMIT
}
