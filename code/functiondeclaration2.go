package main

import "fmt"

// START OMIT
func main() {
	var divident = 503
	divideby := func(d int) (quotient int, mod int) { // HL
		return divident / d, divident % d // HL
	}
	fmt.Println(divideby(2)) // HL
}

// END OMIT
