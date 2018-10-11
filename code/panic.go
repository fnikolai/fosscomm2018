package main

import (
	"fmt"
	"os"
)

func main() {
	// START OMIT
	defer func() { // HL
		msg := recover() // HL
		fmt.Println("Recover from panic errror:", msg)
	}()
	f, err := os.Create("./")
	if err != nil {
		panic(err) // HL
	}
	defer f.Close() // HL
	fmt.Fprintf(f, "hello")
	// END OMIT
}
