package main

import "fmt"

// START OMIT
func main() {
	var divident = 503
	divideby := func(d int) (quotient int, mod int) { // Assign function to value // HL
		return divident / d, divident % d // Lexically scoped name binding // HL
	}
	fmt.Println(divideby(2)) // HL
}

// END OMIT
